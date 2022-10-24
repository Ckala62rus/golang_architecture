1. генерация swagger docs
   swag init -g ./cmd/main.go

2. запуск проекта
   cls; go run .\cmd\main.go

3. сборка проекта
   go build .\cmd\main.go

4. для отладки в VS Code
   {
   // Use IntelliSense to learn about possible attributes.
   // Hover to view descriptions of existing attributes.
   // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
   "version": "0.2.0",
   "configurations": [
   {
   "name": "Launch",
   "type": "go",
   "request": "launch",
   "mode": "auto",
   "cwd": "${workspaceFolder}",
            "program": "${workspaceFolder}/cmd",
   "env": {},
   "args": []
   }
   ]
   }
