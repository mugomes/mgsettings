# MGSettings

[![License](https://img.shields.io/badge/license-PolyForm%20Perimeter%201.0.0-5351FB)](LICENSE.md)

**MGSettings** é uma biblioteca leve em Go para **gerenciamento simples de configurações persistentes**, utilizando arquivos JSON armazenados automaticamente no diretório do usuário ou em um caminho customizado.

Ideal para aplicações CLI, desktop ou serviços que precisam salvar preferências sem dependências externas.

---

## ✨ Recursos

* 💾 Persistência automática em arquivo `config.json`
* 🏠 Suporte a diretório padrão no *Home* do usuário (`~/.appname`)
* 📁 Suporte a caminho customizado
* 🧩 API simples para tipos comuns
  * `string`
  * `int`
  * `bool`
  * `[]string`
* 🔄 Fallback automático para valores padrão
* 📦 Baseado apenas na biblioteca padrão do Go

---

## 📦 Instalação

```bash
go get github.com/profmugomes/mgsettings
```

---

## 🚀 Uso básico

### Carregando configurações

```go
import "github.com/profmugomes/mgsettings"

cfg, err := mgsettings.Load("meuapp", true)
if err != nil {
	log.Fatal(err)
}
```

Isso criará automaticamente:

```text
~/.meuapp/config.json
```

---

## ✍️ Salvando valores

```go
cfg.SetString("username", "joao")
cfg.SetInt("port", 8080)
cfg.SetBool("dark_mode", true)
cfg.SetStringSlice("languages", []string{"pt", "en"})

cfg.Save()
```

---

## 📖 Lendo valores com fallback

```go
user := cfg.GetString("username", "guest")
port := cfg.GetInt("port", 3000)
dark := cfg.GetBool("dark_mode", false)
langs := cfg.GetStringSlice("languages", []string{"en"})
```

Se a chave não existir, o valor padrão será retornado.

---

## 🧠 Como funciona

* As configurações são armazenadas internamente como `json.RawMessage`
* Cada valor é serializado individualmente
* O arquivo só é gravado quando `Save()` é chamado
* Tipos são preservados automaticamente

---

## 🧩 Estrutura do arquivo gerado

```json
{
  "username": "joao",
  "port": 8080,
  "dark_mode": true,
  "languages": [
    "pt",
    "en"
  ]
}
```

---

## 🧩 Compatibilidade

* Go 1.26.5+

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved.

This project is licensed under the PolyForm Perimeter License 1.0.1.

### Summary

This software is available for commercial and noncommercial use, subject to the terms of the PolyForm Perimeter License 1.0.1.

You may:

* ✔ Use the software for commercial and noncommercial purposes.
* ✔ Inspect and study the source code.
* ✔ Modify the software.
* ✔ Create derivative works based on the software.
* ✔ Redistribute the software and permitted modifications.

You may not:

* ✖ Provide a product that competes with the software.

See the full license terms at LICENSE.md.

This summary is provided for convenience only and does not replace or modify the full license terms.