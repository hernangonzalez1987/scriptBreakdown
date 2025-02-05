package ai

import (
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func Test_ai_AnalyzeSceneText(t *testing.T) {

	ctx := context.Background()

	resp := "```json\n{\n  \"props\": [\n    \"mesa de trabalho\",\n    \"computador\",\n    \"mochila\",\n    \"papéis\",\n    \"pastas\"\n  ],\n  \"cast\": [\n    \"Empregado (35, branco)\",\n    \"Gerente (35, branca)\"\n  ],\n  \"wardrobe\": [\n    \"camisa social\"\n  ],\n  \"vehicles\": [],\n  \"animals\": [],\n  \"sound-effects\": [\n    \"luzes apagando\",\n    \"som de passos (sombra na janela)\"\n  ],\n  \"music\": [],\n  \"special-effects\": [\n    \"luzes apagando\",\n    \"reflexo de um homem de máscara na tela do computador\"\n  ],\n  \"set-dressing\": [\n    \"sala de escritório moderna e bem decorada\",\n    \"escrivaninhas vazias\",\n    \"janela com vista para a cidade noturna\"\n  ]\n}\n```\n"

	model := NewFakeLLM([]string{
		resp,
	})

	cache := _mocks.NewMockCache(t)
	cache.EXPECT().Get("509b423e-1afb-5e5a-9f90-4691db481dec").Return("", false)
	cache.EXPECT().Save("509b423e-1afb-5e5a-9f90-4691db481dec", resp)

	ai := ai{model: model, cache: cache}

	scene := "Uma mesa de trabalho com um computador ligado em uma ampla sala de escritório moderna e bem decorada. \nUm EMPREGADO (35, branco), de camisa social, trabalha focado na tela. Todas as escrivaninhas próximas estão vazias. Pela janela, a noite já caiu sobre a cidade.\nA GERENTE (35, branca), a única outra pessoa no lugar, se aproxima.\nGERENTE\nOpa, já estou saindo. Não quer deixar para terminar isso amanhã?\nEMPREGADO\nTô quase acabando, aqui.\nGERENTE\nVocê desliga as luzes ao sair?\nEMPREGADO\nPode deixar.\nA gerente vai embora. O empregado olha o relógio e continua trabalhando.\nDe repente, as LUZES APAGAM.\nEMPREGADO\nÔ! Ainda tô aqui! \nNinguém responde. O homem, agora iluminado apenas pela tela do computador, volta o seu olhar para o texto.\nEMPREGADO\nFilha da…\nUma sombra passa perto da janela. Tensão.\nEMPREGADO\nSílvia?\nO empregado olha ao redor, não tem ninguém. Olha o relógio mais uma vez.\nNão adianta continuar trabalhando assim. Junta as suas coisas, guarda alguns papéis e pastas numa mochila. \nDesliga o computador. De repente, na tela preta se reflete um homem de máscara detrás do empregado."

	tags, err := ai.AnalyzeSceneText(ctx, scene)

	expected := []entity.Tag{
		{Category: "set-dressing", Element: "sala de escritório moderna e bem decorada", Description: ""},
		{Category: "set-dressing", Element: "escrivaninhas vazias", Description: ""},
		{Category: "set-dressing", Element: "janela com vista para a cidade noturna", Description: ""},
		{Category: "sound-effects", Element: "luzes apagando", Description: ""},
		{Category: "sound-effects", Element: "som de passos (sombra na janela)", Description: ""},
		{Category: "wardrobe", Element: "camisa social", Description: ""},
		{Category: "special-effects", Element: "luzes apagando", Description: ""},
		{Category: "special-effects", Element: "reflexo de um homem de máscara na tela do computador", Description: ""},
		{Category: "props", Element: "mesa de trabalho", Description: ""},
		{Category: "props", Element: "computador", Description: ""},
		{Category: "props", Element: "mochila", Description: ""},
		{Category: "props", Element: "papéis", Description: ""},
		{Category: "props", Element: "pastas", Description: ""},
		{Category: "cast", Element: "Empregado (35, branco)", Description: ""},
		{Category: "cast", Element: "Gerente (35, branca)", Description: ""},
	}

	assert.NoError(t, err)
	assert.ElementsMatch(t, expected, tags)

}

func Test_parseResponse(t *testing.T) {

	resp := "```json\n{\n  \"props\": [\n    \"mesa de trabalho\",\n    \"computador\",\n    \"mochila\",\n    \"papéis\",\n    \"pastas\"\n  ],\n  \"cast\": [\n    \"Empregado (35, branco)\",\n    \"Gerente (35, branca)\"\n  ],\n  \"wardrobe\": [\n    \"camisa social\"\n  ],\n  \"vehicles\": [],\n  \"animals\": [],\n  \"sound-effects\": [\n    \"luzes apagando\",\n    \"som de passos (sombra na janela)\"\n  ],\n  \"music\": [],\n  \"special-effects\": [\n    \"luzes apagando\",\n    \"reflexo de um homem de máscara na tela do computador\"\n  ],\n  \"set-dressing\": [\n    \"sala de escritório moderna e bem decorada\",\n    \"escrivaninhas vazias\",\n    \"janela com vista para a cidade noturna\"\n  ]\n}\n```\n"

	parsed, err := parseResponse(resp)

	expected := map[string][]string{
		"wardrobe":        {"camisa social"},
		"music":           {},
		"special-effects": {"luzes apagando", "reflexo de um homem de máscara na tela do computador"},
		"set-dressing":    {"sala de escritório moderna e bem decorada", "escrivaninhas vazias", "janela com vista para a cidade noturna"},
		"props":           {"mesa de trabalho", "computador", "mochila", "papéis", "pastas"},
		"cast":            {"Empregado (35, branco)", "Gerente (35, branca)"},
		"vehicles":        {},
		"animals":         {},
		"sound-effects":   {"luzes apagando", "som de passos (sombra na janela)"},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, parsed)
}

func Test_extractTagsFromResponse(t *testing.T) {

	parsed := map[string][]string{
		"wardrobe":        {"camisa social"},
		"music":           {},
		"special-effects": {"luzes apagando", "reflexo de um homem de máscara na tela do computador"},
		"set-dressing":    {"sala de escritório moderna e bem decorada", "escrivaninhas vazias", "janela com vista para a cidade noturna"},
		"props":           {"mesa de trabalho", "computador", "mochila", "papéis", "pastas"},
		"cast":            {"Empregado (35, branco)", "Gerente (35, branca)"},
		"vehicles":        {},
		"animals":         {},
		"sound-effects":   {"luzes apagando", "som de passos (sombra na janela)"},
	}

	expected := []entity.Tag{
		{Category: "set-dressing", Element: "sala de escritório moderna e bem decorada", Description: ""},
		{Category: "set-dressing", Element: "escrivaninhas vazias", Description: ""},
		{Category: "set-dressing", Element: "janela com vista para a cidade noturna", Description: ""},
		{Category: "sound-effects", Element: "luzes apagando", Description: ""},
		{Category: "sound-effects", Element: "som de passos (sombra na janela)", Description: ""},
		{Category: "wardrobe", Element: "camisa social", Description: ""},
		{Category: "special-effects", Element: "luzes apagando", Description: ""},
		{Category: "special-effects", Element: "reflexo de um homem de máscara na tela do computador", Description: ""},
		{Category: "props", Element: "mesa de trabalho", Description: ""},
		{Category: "props", Element: "computador", Description: ""},
		{Category: "props", Element: "mochila", Description: ""},
		{Category: "props", Element: "papéis", Description: ""},
		{Category: "props", Element: "pastas", Description: ""},
		{Category: "cast", Element: "Empregado (35, branco)", Description: ""},
		{Category: "cast", Element: "Gerente (35, branca)", Description: ""},
	}

	tags := extractTagsFromResponse(parsed)

	assert.ElementsMatch(t, expected, tags)

}
