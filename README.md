# os-json-to-cvs-uncompress

A ideia é partir de um arquivo json, gerar um arquivo csv.

Como o arquivo json gerado por meio de dump do KVS gera um Object Store do dynamoDB, o campo com valores vem compactado.

Então, esse projeto descompacta esse valor que vem em 'compressed_value'