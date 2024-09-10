# cobra-cli-samples

这个是一个使用Go语言编写的CLI应用程序示例，它使用了Cobra和Viper库。该应用程序提供了一系列命令（通过Cobra实现），用于使用Viper从配置文件中读取、写入、更新和删除配置记录。
## Using the example CLI App.

`<cli> config -h` 提供了以下文档，用于使用这个CLI应用程序

```
可用命令： add 'add' 子命令将传递的键值对添加到应用程序配置文件中。
 delete 'delete' 子命令从配置文件中移除键值对。 
 update 'update' 子命令将传递的键值对更新到现有的数据集中到应用程序配置文件。
  view 'view' 子命令将提供键的列表和值的映射。
  
标志： -h, --help config的帮助信息 
-k, --key string 要添加到配置中的键值对的键。 
-v, --value string 要添加到配置中的键值对的值。
```


### Examples, The CRUD!

./cli config add -k "blog" -v "https://compositecode.blog/" 示例将一条记录写入配置文件，键为 "blog"，值为 "https://compositecode.blog/"。
./cli config view 显示配置文件的内容和CLI特定的环境变量。这些配置文件位于 .cobrae-cli-samples.yml 中，环境变量以 COBRACLISAMPLES 开头。首先显示键的列表，然后在其下方显示键和值。
./cli config update -k "blog" -v "not found" 将更新配置中的博客条目，将其值更改为 "not found"。
./cli config delete ... 将从配置文件中删除键和值。

## Building the Project
go test
go build -o cli_demo


## 加入全局变量 所有用户生效

sudo sh -c 'echo "export PATH=\$PATH:/opt/cli_demo*" >> /etc/profile && echo "alias wws=\"/opt/cli_demo*\"" >> /etc/profile'

source /etc/profile


