# jupyter 安装并配置 go

安装 Jupyter notebook

```sh
# 最合适，每个工具独立虚拟环境，自带 Python 版本管理，卸载干净快速
# 本质上给工具单独开一个 venv，入口命令软链到 ~/.local/bin
# 这样安装会保留 jupyter 和 jupyter-lab 命令
uv tool install jupyterlab --with-executables-from jupyter-core --force

# formula 会拖一个 brew 版本的 Python 进来
# 不建议用 brew 安装 Python，容易出问题，brew 更适合安装系统级二进制工具
brew install jupyterlab

# 经常切换版本不建议
pip3 install jupyterlab

# VSCode 插件安装，依赖 Apple 命令行工具自带的 Python，版本很老
# 且 Xcode 更新可能会换掉，包括路径也更换，导致安装的包也一起失效
```

安装内核

```sh
# 查看当前工具
uv tool list

# 查看可用内核
jupyter kernelspec list

# go 和 gopls 已经安装了，只需要安装 gonb goimports
go install github.com/janpfeifer/gonb@latest && go install golang.org/x/tools/cmd/goimports@latest
gonb --install

# 顺便安装一下 deno 和内核
brew install deno
deno jupyter --install
```

go 要在 Jupyter notebook 上使用时，要安装 gonb
gonb 会把每个 cell 包进要给隐式的 main 函数执行，所以直接写语句就能执行
要定义包级的东西（func type var const）它会自动识别并提示到包级，跨 cell 保留
`%%` 开头是 gonb 的魔法命令

初学者学习 包/模块/init 顺序等内容时，要结合多文件、多目录的概念，放在 notbook 的 cell 里反而看不出来，学习语法胡总和标准库 api 的时候才适合使用 gonb 快速查看结果

ctrl enter 运行，shift enter 运行并创建下一块，alt enter 强制执行并创建下一块，tab 提示代码

魔法符号会自动导入，自动包裹 main 方法

***工欲善其事，必先利其器***
***但是切记，不要在学习的过程中突然沉迷于工具配置了！***
