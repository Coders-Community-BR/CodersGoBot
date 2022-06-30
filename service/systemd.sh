#!/bin/bash

service_path="/etc/systemd/system"
ln_service_path="lib/systemd/system"
bin_path="/root/go/bin/CodersGoBot"

if [ "$EUID" -ne 0 ]
then 
    echo "Por favor, rode o script como root."
    echo "Saindo..."
    exit 1
fi

if ! [ -x "$(command -v go)" ];
then
    echo "Parece que você não possui Go instalado no sistema. Go é necessário para rodar o nosso bot, então por favor, primeiro instale ele e depois tente executar o script novamente."
    echo "Saindo..."
    exit 1
fi

if [[ "$(pwd)" == *service ]]
then
    cd ..
fi

if [ ! -d "$bin_path" ];
then
    mkdir /root/go/bin/CodersGoBot
fi

go build -o "${bin_path}/bot"
cp ./.env "${bin_path}"
echo "Build efetuada com sucesso."
cp ./service/codersbot.service "${service_path}"

systemctl start codersbot.service
systemctl enable codersbot.service
if [ ! -f "${ln_service_path}/codersbot.service" ];
then
    ln -s /etc/systemd/system/codersbot.service /lib/systemd/system
fi
echo "Serviço criado e iniciado."