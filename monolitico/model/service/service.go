package service

import (
	"log"
	"errors"
	"github.com/Pedro-Bezerra/reservalab-mono/database"
	"github.com/Pedro-Bezerra/reservalab-mono/model/dto"
	"github.com/Pedro-Bezerra/reservalab-mono/model/entity"
)

// FALTA ADICIONAR NO BANCO DE DADOS

func EnviarMensagem(mensagemDto *dto.MensagemDto) (*entity.Mensagem, error) {

	var mensagem entity.Mensagem = *dto.CriarMensagem(mensagemDto)

	if err := database.DB.Create(&mensagem).Error; err != nil {
		log.Fatalf("Erro no armazenamento da mensagem: %v", err)
		return nil, err
	}

	return &mensagem, nil
}

func EnviarSolicitacao(solicitacaoDto *dto.SolicitacaoDto) (*entity.Solicitacao, error) {

	var solicitacao entity.Solicitacao = *dto.CriarSolicitacao(solicitacaoDto)

	isDisponivel, erro := ChecarDisponibilidade(solicitacaoDto)

	if !isDisponivel {
		return nil, erro
	}

	if err := database.DB.Create(&solicitacao).Error; err != nil {
		log.Fatalf("Erro no armazenamento da solicitação: %v", err)
		return nil, err
	}

	reserva := entity.CriarReserva(&solicitacao)

	if err := database.DB.Create(&reserva).Error; err != nil {
		return nil, err
	}



	return &solicitacao, nil
}

func ChecarDisponibilidade(solicitacaoDto *dto.SolicitacaoDto) (bool, error) {

	var reservas []entity.Reserva

	if err := database.DB.Where("data = ? AND horario_inicio < ? AND horario_termino > ?", solicitacaoDto.Data, solicitacaoDto.HorarioTermino, solicitacaoDto.HorarioInicio).Find(&reservas).Error; err != nil {
		log.Fatalf("Erro na obtenção das reservas: %v", err)
		return false, err
	}

	if len(reservas) > 0 {
		return false, errors.New("HÁ OUTRA RESERVA NA DATA E NO HORÁRIO INDICADOS")
	}

	return true, nil
}

func AtualizarSolicitacao(idSolicitacao uint, solicitacaoNova *dto.SolicitacaoDto) (*entity.Solicitacao, error) {

	isDisponivel, erroDisponibilidade := ChecarDisponibilidade(solicitacaoNova)

	if !isDisponivel {
		return nil, erroDisponibilidade
	}

	solicitacaoAtualizar := entity.Solicitacao{
		IdSolicitacao: idSolicitacao,
		Nome: solicitacaoNova.Nome,
		Email: solicitacaoNova.Email,
		Matricula: solicitacaoNova.Matricula,
		Vinculo: solicitacaoNova.Vinculo,
		Observacoes: solicitacaoNova.Observacoes,
		Data: solicitacaoNova.Data,
		HorarioInicio: solicitacaoNova.HorarioInicio,
		HorarioTermino: solicitacaoNova.HorarioTermino,
	}

	_, erroDelete := RemoverReserva(idSolicitacao)

	if erroDelete != nil {
		return nil, erroDelete
	}

	if err := database.DB.Save(&solicitacaoAtualizar).Error; err != nil {
		return nil, errors.New("ERRO NO REAGENDAMENTO DA RESERVA")
	}

	reserva := entity.CriarReserva(&solicitacaoAtualizar)

	if err := database.DB.Create(&reserva).Error; err != nil {
		return nil, err
	}

	return &solicitacaoAtualizar, nil


}

func GetSolicitacaoByMatricula(matricula string) (*[]entity.Solicitacao, error) {
	var solicitacoes []entity.Solicitacao

	if err := database.DB.Where("matricula = ?", matricula).Find(&solicitacoes).Error; err != nil {
		log.Fatalf("Erro na obtenção das solicitações: %v", err)
		return nil, err
	}

	if len(solicitacoes) == 0 {
		return nil, errors.New("NÃO HÁ SOLICITAÇÕES REALIZADAS USANDO A MATRÍCULA INDICADA")
	}

	return &solicitacoes, nil
}

func RemoverReserva(fk uint) (bool, error) {

	if err := database.DB.Where("fk_solicitacao = ?", fk).Delete(&entity.Reserva{}).Error; err != nil {
		return false, errors.New("ERRO NA ELIMINAÇÃO DA RESERVA")
	}

	return true, nil

}

