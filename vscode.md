
vscode setting.json

```
 "workbench.startupEditor": "newUntitledFile",
  // 控制字体系列。
  "editor.fontFamily": "Consolas, 'Courier New', monospace",
  // 启用字体连字
  "editor.fontLigatures": false,
  // 控制字体粗细。
  "editor.fontWeight": "normal",
  // 以像素为单位控制字号。
  "editor.fontSize": 14,
  "editor.tabSize": 4,
  "files.exclude": {
    "**/.DS_Store": true,
    "**/.git": true,
    "**/.hg": true,
    "**/.idea": true,
    "**/.svn": true,
    "**/build": true,
    "**/CVS": true,
    "**/dist": true,
    "**/node_modules": true,
  },
  "go.formatTool": "goimports",
  "git.enableSmartCommit": true,
  "diffEditor.ignoreTrimWhitespace": false,
  "go.languageServerExperimentalFeatures": {
    "diagnostics": true, // 提供 build 和 vet 的报错信息
    "documentLink": true // import 可以跳转到项目的 godoc (前提是已经在公有仓库里并且被收录)
  },
  "editor.codeActionsOnSave": {
    "source.organizeImports": true
  },
  "editor.snippetSuggestions": "none",
  "editor.formatOnSave": true,
  "gopls": {
    "usePlaceholders": true, // 在完成一个函数时添加参数占位符
    "completionDocumentation": true // 用于完成项中的文档
  },
  "mssql.connections": [
    {
      "server": "{{put-server-name-here}}",
      "database": "{{put-database-name-here}}",
      "user": "{{put-username-here}}",
      "password": "{{put-password-here}}"
    }
  ],
  "go.toolsEnvVars": {
    "GOFLAGS": "-mod=vendor"
  },
  "go.useLanguageServer": false,
  "[jsonc]": {
    "editor.quickSuggestions": {
      "strings": true
    },
    "editor.suggest.insertMode": "replace"
  },
  "update.mode": "none",
  "C_Cpp.updateChannel": "Insiders",
```