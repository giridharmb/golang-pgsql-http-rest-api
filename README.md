#### HTTP REST API & Query PGSQL Backend

<b>This program connects to PGSQL database and makes a query <br/>
against a table and fetches all the rows. This is exposed via <br/>
HTTP REST API (GET) call to fetch all the rows as a JSON</b>

Please refer to [HTTP Server And REST API](#https://github.com/giridharmb/golang-http-server-rest-api) <br/>
For setting up HTTP Server/REST API.

Please refer to [This Document To Generate Sample PostgreSQL Table](https://github.com/giridharmb/PostgreSQL-README#generate-a-random-table) <br/>
Before running the below code.

#### Fetch The Dependencies

```bash
go mod init; go mod tidy;
```

```bash
go get -u -v ./...
```

#### Build The Binary

```bash
go build
```

#### Run The Program

```bash
./golang-pgsql-http-rest-api
```

#### Make HTTP GET Request To Fetch DB Records from PostgreSQL

```bash
curl \
-H "accept:application/json" \
-H "content-type:application/json" \
-X GET http://localhost:9000/api/v1/getDataFromPGSQL 2>/dev/null | python -m json.tool
```

#### Output

```json
{
    "error": "",
    "serverResponse": [
        {
            "md5": "4148034f7ff59deddce0507cf5df6136",
            "random_float": 0.04866303424917717,
            "random_num": 642
        },
        {
            "md5": "50cf40c1cd51fa0c9bcda04d19118637",
            "random_float": 0.42088215715087784,
            "random_num": 299
        },
        {
            "md5": "d83b110f8c24dc876feb39e1d32fc9a9",
            "random_float": 0.5040329323515316,
            "random_num": 452
        },
        {
            "md5": "bd5e6593d1133df5abbe2dcd7ed41b22",
            "random_float": 0.680895277194864,
            "random_num": 558
        },
        {
            "md5": "b0e288ee1d4d5ea7c50995091222526c",
            "random_float": 0.9993317095718943,
            "random_num": 436
        },
        {
            "md5": "1ecf469484a66c646f33f1ad26d2c120",
            "random_float": 0.4972750511722168,
            "random_num": 143
        },
        {
            "md5": "04705f29ec44fede34940cd954acf20c",
            "random_float": 0.7588865512592768,
            "random_num": 987
        },
        {
            "md5": "fe0f5b228256a75beccfe7156af05616",
            "random_float": 0.9723930089149349,
            "random_num": 337
        },
        {
            "md5": "5522e09f9a97abd695bf821ce1eca750",
            "random_float": 0.637456665167111,
            "random_num": 321
        },
        {
            "md5": "34726f12c5bae77be4b3bb0293ccb38a",
            "random_float": 0.9731822082784944,
            "random_num": 896
        },
        {
            "md5": "72a24d55ebdc63031fd1b1bb0b4b20a4",
            "random_float": 0.21032232956669006,
            "random_num": 647
        },
        {
            "md5": "03d076d7a4477f7b7800572cfed4678d",
            "random_float": 0.7768539272009178,
            "random_num": 838
        },
        {
            "md5": "88b31b9ca28b6977bc128a1e5ec0bceb",
            "random_float": 0.923863336295458,
            "random_num": 195
        },
        {
            "md5": "7042290814b38099ce380b75343fde1f",
            "random_float": 0.9273458418532314,
            "random_num": 641
        },
        {
            "md5": "45cd6c05e25c992f909c2cb05a3d2a8f",
            "random_float": 0.8078386414598491,
            "random_num": 806
        },
        {
            "md5": "37804d9a3a95c24abdc819d7a30ec59a",
            "random_float": 0.46808880132621766,
            "random_num": 192
        },
        {
            "md5": "eaa7062ce381b59a5c932b1e9751a6d4",
            "random_float": 0.2319390696668684,
            "random_num": 215
        },
        {
            "md5": "03983dceb4553341d65ba64dc652a810",
            "random_float": 0.42519699483093376,
            "random_num": 974
        },
        {
            "md5": "3c560e0683197ec18af06894d3f83301",
            "random_float": 0.41466590628031597,
            "random_num": 271
        },
        {
            "md5": "2ed21a8da890b55f0c093e6297b461dd",
            "random_float": 0.4320334559281065,
            "random_num": 885
        },
        {
            "md5": "3f829ee41011d1491131fcc71da9265f",
            "random_float": 0.9617127469236983,
            "random_num": 246
        },
        {
            "md5": "9b96a2e7bb9c6affe66ba39b3e70711e",
            "random_float": 0.49537617767868625,
            "random_num": 387
        },
        {
            "md5": "6514fa379678e067b6e4d3542fd07663",
            "random_float": 0.007216044821998224,
            "random_num": 681
        },
        {
            "md5": "1748c24b00f76dcd18147d09521218d6",
            "random_float": 0.11057322298324834,
            "random_num": 486
        },
        {
            "md5": "7e97ec347cef5603a50005ae5e4b4676",
            "random_float": 0.3100045958117228,
            "random_num": 626
        },
        {
            "md5": "3a4df72e8a5a216af3c482d95dd20257",
            "random_float": 0.7408985296407238,
            "random_num": 825
        },
        {
            "md5": "e8d18c6f96c064ac0ed9dc262bbb374c",
            "random_float": 0.553555514391018,
            "random_num": 576
        },
        {
            "md5": "af1ddfb222435fb90ae6a645580c3caf",
            "random_float": 0.7783012875830053,
            "random_num": 888
        },
        {
            "md5": "9ab3df55e882dbf14329d5e530265ce0",
            "random_float": 0.19930670196319866,
            "random_num": 764
        },
        {
            "md5": "674bc7031858e7e6223c24bf914b07fc",
            "random_float": 0.15416282690015493,
            "random_num": 954
        },
        {
            "md5": "5d99c44375a838868d7d61b5ecdfe003",
            "random_float": 0.20999317590922928,
            "random_num": 124
        },
        {
            "md5": "92234780cbeff32bd4a84b077e337471",
            "random_float": 0.614202380679302,
            "random_num": 665
        },
        {
            "md5": "5461bbba0ef0b60a784dc98017c5f963",
            "random_float": 0.28775154207751896,
            "random_num": 294
        },
        {
            "md5": "ac019f632d594763a94ec3235fb0709d",
            "random_float": 0.3771571342124034,
            "random_num": 969
        },
        {
            "md5": "926ecf93c031479d71652d1110c69795",
            "random_float": 0.6073090913127253,
            "random_num": 935
        },
        {
            "md5": "7ebaa3c1d51637612f9671cfe306ce1d",
            "random_float": 0.15174146166908997,
            "random_num": 808
        },
        {
            "md5": "f1075e715496b08ac6f544fa1ad52919",
            "random_float": 0.06329444449005805,
            "random_num": 931
        },
        {
            "md5": "f66e350fcd57a4e48539436b6ff853f5",
            "random_float": 0.4118590686840733,
            "random_num": 790
        },
        {
            "md5": "0552f48669c338021ca4f245fefcf625",
            "random_float": 0.3780827530611326,
            "random_num": 800
        },
        {
            "md5": "621901a283a8da41581c8a3d2b979f0f",
            "random_float": 0.9468833288514702,
            "random_num": 560
        },
        {
            "md5": "43d9d58ba18655064de813ce708e695a",
            "random_float": 0.7485215581686369,
            "random_num": 799
        },
        {
            "md5": "f575436dec30eea1bab4756f19fbf824",
            "random_float": 0.7825670782088352,
            "random_num": 980
        },
        {
            "md5": "9923f003b7f65d073854f90c9324972b",
            "random_float": 0.3338920222316908,
            "random_num": 576
        },
        {
            "md5": "83e0068c7a176a7d129de374557f4f17",
            "random_float": 0.3029625589262963,
            "random_num": 332
        },
        {
            "md5": "a925dedc99554cf6479a737202a6d7b5",
            "random_float": 0.6991179885696148,
            "random_num": 321
        },
        {
            "md5": "8fa485bb433aa92c4084fdc803aab0a0",
            "random_float": 0.32020739512856977,
            "random_num": 811
        },
        {
            "md5": "65999c909fa29a893483d3213ad2fd71",
            "random_float": 0.7001050099240693,
            "random_num": 406
        },
        {
            "md5": "e7e637b22fec67fc227df18f55a80bfb",
            "random_float": 0.5647278703065979,
            "random_num": 987
        },
        {
            "md5": "8361969cbf9834968e7a7c5dcd30620f",
            "random_float": 0.11918655102194364,
            "random_num": 636
        },
        {
            "md5": "6fedebc4f5861884eb32b98e45158838",
            "random_float": 0.06389097046339742,
            "random_num": 943
        },
        {
            "md5": "4df398ddda29b054276e2c5bf80fbd04",
            "random_float": 0.9052090910475776,
            "random_num": 956
        },
        {
            "md5": "c66d0da73e34d335ff3458576b5bce1a",
            "random_float": 0.943839991589428,
            "random_num": 246
        },
        {
            "md5": "fa9982a24e148ddf150121fcfadebb4d",
            "random_float": 0.7998664118875958,
            "random_num": 815
        },
        {
            "md5": "4035f004b0e0a7ac3d3518564f9bd461",
            "random_float": 0.5393722880483622,
            "random_num": 982
        },
        {
            "md5": "d9c1d80a74dc06b850dff69f2748973e",
            "random_float": 0.6841220025018764,
            "random_num": 180
        },
        {
            "md5": "11d9c1525cf6a74dd669e58c945f05d9",
            "random_float": 0.08902145857075539,
            "random_num": 104
        },
        {
            "md5": "d311a1b52a5850f90544c56614de78d6",
            "random_float": 0.633008047170609,
            "random_num": 781
        },
        {
            "md5": "90f348bd16902118839424dc07b01410",
            "random_float": 0.7066890636274294,
            "random_num": 857
        },
        {
            "md5": "d274c8ed6958c6a767d211b8dd012a95",
            "random_float": 0.7580923617955015,
            "random_num": 400
        },
        {
            "md5": "4e2681e44b845aee47f301b492f4b0af",
            "random_float": 0.8622295231997867,
            "random_num": 367
        },
        {
            "md5": "902793e92bce6a2bed53f7d8c7858669",
            "random_float": 0.3873367600145663,
            "random_num": 995
        },
        {
            "md5": "4f5092ce9d22f9b4fa4ffd4858e1423d",
            "random_float": 0.7961150138870963,
            "random_num": 177
        },
        {
            "md5": "888ff9f3e78f8533ec2b70c1c8a9c6ef",
            "random_float": 0.39937875977346593,
            "random_num": 942
        },
        {
            "md5": "197c5309bd05006232712ac2ae56a63b",
            "random_float": 0.6458135983466207,
            "random_num": 656
        },
        {
            "md5": "b63816fb4dd6ff132bd455d500527791",
            "random_float": 0.5482767751396587,
            "random_num": 253
        },
        {
            "md5": "81e3a1a290e924ade15ddc4c51caa479",
            "random_float": 0.7454886970623065,
            "random_num": 619
        },
        {
            "md5": "c5273fe0b05fa5ccfb279ec86dfe7f18",
            "random_float": 0.5630306448459059,
            "random_num": 454
        },
        {
            "md5": "71d7943c7fbcc88eeeeeb203756b09b0",
            "random_float": 0.2746654457044002,
            "random_num": 868
        },
        {
            "md5": "18e01110720c4a3b71fd149a88f7b9fe",
            "random_float": 0.4114819822255562,
            "random_num": 651
        },
        {
            "md5": "38084e856a1a5eb5705255e3806c9db3",
            "random_float": 0.9827164426164998,
            "random_num": 533
        },
        {
            "md5": "a2bd40112c5c232bfbddedc5ae903319",
            "random_float": 0.4912883384241944,
            "random_num": 562
        },
        {
            "md5": "82e067fe6191f0b7c8546228285143cb",
            "random_float": 0.5737877030631466,
            "random_num": 748
        },
        {
            "md5": "990b0807606e9deca40a95d0383607aa",
            "random_float": 0.4405730440451414,
            "random_num": 198
        },
        {
            "md5": "93b6d57d8ef03ee4580f61da824011ec",
            "random_float": 0.381388165085518,
            "random_num": 193
        },
        {
            "md5": "066c0064dd0793e1cc16834118b5ff91",
            "random_float": 0.2732565146791437,
            "random_num": 917
        },
        {
            "md5": "fa916455a897533daebc3fa6fb8fe8a9",
            "random_float": 0.8692834040772226,
            "random_num": 144
        },
        {
            "md5": "5d64ceec3ff1128abeac2973acdfd473",
            "random_float": 0.6174341312967364,
            "random_num": 288
        },
        {
            "md5": "5268e20f66183a5eedd79416690df8c6",
            "random_float": 0.6684679589817293,
            "random_num": 190
        },
        {
            "md5": "865997ac405530db6a8fb0380e2bab70",
            "random_float": 0.512843154737876,
            "random_num": 882
        },
        {
            "md5": "13118ebcc88e61d06b80df1876fd4e61",
            "random_float": 0.5351699310018887,
            "random_num": 169
        },
        {
            "md5": "46dbaaded8984af9e5c9c9fdb4c9db5a",
            "random_float": 0.9538713070687841,
            "random_num": 860
        },
        {
            "md5": "c865d4c87c5b3a53834c95fdbeb829e6",
            "random_float": 0.7412783601801252,
            "random_num": 515
        },
        {
            "md5": "9cb7ef3f26706550bf9ad240e9f6b700",
            "random_float": 0.24512715853164124,
            "random_num": 640
        },
        {
            "md5": "c177af52f089f6f8937eba1339182f3a",
            "random_float": 0.11718442119668282,
            "random_num": 780
        },
        {
            "md5": "5ee0fbae1528416760100a2849b3dae1",
            "random_float": 0.9310488722010213,
            "random_num": 983
        },
        {
            "md5": "f6679ef63da27b4f946888ab94f07026",
            "random_float": 0.913714596952282,
            "random_num": 297
        },
        {
            "md5": "246caa4ae3d071b684955819362c3ee3",
            "random_float": 0.5033302791963159,
            "random_num": 500
        },
        {
            "md5": "39f449cbe253e8db2119ab4cf25a15d5",
            "random_float": 0.2372880839843674,
            "random_num": 467
        },
        {
            "md5": "27e4e24cb71461b5b76d795f3b4b129f",
            "random_float": 0.6818451109800101,
            "random_num": 696
        },
        {
            "md5": "910366af12f237e0c9282a896a9db4b8",
            "random_float": 0.2629046032769722,
            "random_num": 789
        },
        {
            "md5": "49096b5c1b032dfb82d5c94ad7569bcc",
            "random_float": 0.4178056477105194,
            "random_num": 901
        },
        {
            "md5": "9eff0e237642b1e83f7ba0dae88fd649",
            "random_float": 0.9001832382139554,
            "random_num": 243
        },
        {
            "md5": "b304aa397aca27217c399598c3a51f23",
            "random_float": 0.6344251571244719,
            "random_num": 801
        },
        {
            "md5": "c8fadd9ef98eba6c2a9f54b598f7c024",
            "random_float": 0.98834859073785,
            "random_num": 579
        },
        {
            "md5": "ce7adee0a27fee6a5003e23500271d42",
            "random_float": 0.6016736918332377,
            "random_num": 487
        },
        {
            "md5": "39147b0af08cbba4e1c479ef01a5b0cc",
            "random_float": 0.9889932151048413,
            "random_num": 787
        },
        {
            "md5": "f4464d38b4dbc57d27280a27120757fe",
            "random_float": 0.8556114453771961,
            "random_num": 808
        },
        {
            "md5": "4dc9c03e1e733801071ef658a1847bfb",
            "random_float": 0.07378013571902642,
            "random_num": 414
        },
        {
            "md5": "257332bb2ab34d5ec578efbfeefca7fb",
            "random_float": 0.5060695358735465,
            "random_num": 340
        },
        {
            "md5": "e4453fd97e11518f5d45bf455541bb70",
            "random_float": 0.316360648146329,
            "random_num": 470
        }
    ]
}
```