package sql

// InsertMockPokemonWeak insert pokemonweak
var InsertMockPokemonWeak = `
INSERT INTO "pokemon_weaks"("pokemon_id","attribute") VALUES
('54a7c0e7-885f-4e8e-92fe-59bd8ddadcc2','みず'),
('54a7c0e7-885f-4e8e-92fe-59bd8ddadcc2','でんき'),
('54a7c0e7-885f-4e8e-92fe-59bd8ddadcc2','いわ'),
('80b694ab-e49c-4d11-800c-ce1deed87a82','ほのお'),
('80b694ab-e49c-4d11-800c-ce1deed87a82','でんき'),
('80b694ab-e49c-4d11-800c-ce1deed87a82','こおり'),
('80b694ab-e49c-4d11-800c-ce1deed87a82','ひこう'),
('80b694ab-e49c-4d11-800c-ce1deed87a82','いわ'),
('88d7d77b-1373-49ff-ba18-513fb1375de0','どく'),
('88d7d77b-1373-49ff-ba18-513fb1375de0','はがね'),
('e9afd9cd-2b49-4462-869e-20b1b3101154','みず'),
('e9afd9cd-2b49-4462-869e-20b1b3101154','じめん'),
('e9afd9cd-2b49-4462-869e-20b1b3101154','いわ'),
('89240025-dc87-4b7e-a33b-b180b9a6bac6','ほのお'),
('89240025-dc87-4b7e-a33b-b180b9a6bac6','こおり'),
('89240025-dc87-4b7e-a33b-b180b9a6bac6','ひこう'),
('89240025-dc87-4b7e-a33b-b180b9a6bac6','エスパー'),
('1bf84a8d-2a0a-4369-b8ca-07bf52c8d03b','みず'),
('1bf84a8d-2a0a-4369-b8ca-07bf52c8d03b','じめん'),
('1bf84a8d-2a0a-4369-b8ca-07bf52c8d03b','いわ'),
('ed822fd9-b396-441e-a450-19567e4b417f','どく'),
('ed822fd9-b396-441e-a450-19567e4b417f','ゴースト'),
('ed822fd9-b396-441e-a450-19567e4b417f','はがね'),
('cca2247d-39e4-4710-a7af-886f80f3122c','くさ'),
('cca2247d-39e4-4710-a7af-886f80f3122c','でんき'),
('cca2247d-39e4-4710-a7af-886f80f3122c','かくとう'),
('cca2247d-39e4-4710-a7af-886f80f3122c','いわ'),
('5e87c549-ad28-4b52-a406-9d6c68d92b2c','じめん'),
('5e87c549-ad28-4b52-a406-9d6c68d92b2c','エスパー'),
('5e87c549-ad28-4b52-a406-9d6c68d92b2c','ゴースト'),
('5e87c549-ad28-4b52-a406-9d6c68d92b2c','あく'),
('54b9fc4e-840b-4e05-ae9c-d2be2b44e984','じめん'),
('54b9fc4e-840b-4e05-ae9c-d2be2b44e984','エスパー'),
('54b9fc4e-840b-4e05-ae9c-d2be2b44e984','はがね'),
('90a8d1d6-0e0d-4a9f-a63a-88c1e44eb74f','でんき'),
('90a8d1d6-0e0d-4a9f-a63a-88c1e44eb74f','いわ'),
('5691d340-6147-450d-bb80-d44f644714ad','くさ'),
('5691d340-6147-450d-bb80-d44f644714ad','でんき'),
('608ae62a-477f-45f8-ac82-735d5063337b','じめん'),
('4afa9e91-61a1-4f1c-9da9-f96468109dc3','みず'),
('4afa9e91-61a1-4f1c-9da9-f96468109dc3','じめん'),
('4afa9e91-61a1-4f1c-9da9-f96468109dc3','いわ'),
('46df8cf3-2b33-477c-b8e7-143ea975078f','かくとう'),
('6574df46-d766-4be0-a79c-f10f8bae8367','くさ'),
('6574df46-d766-4be0-a79c-f10f8bae8367','じめん'),
('7d089493-c64e-4fa4-92c6-19034c0cf310','くさ'),
('22a307c2-1874-4afc-b18c-4b4a9a11a246','むし'),
('22a307c2-1874-4afc-b18c-4b4a9a11a246','ゴースト'),
('22a307c2-1874-4afc-b18c-4b4a9a11a246','あく'),
('ace97de5-e855-42ae-96b7-4c68d5dcb99b','かくとう'),
('ace97de5-e855-42ae-96b7-4c68d5dcb99b','むし'),
('ace97de5-e855-42ae-96b7-4c68d5dcb99b','フェアリー'),
('1b72d661-b90f-4e67-853a-4cf428cc80fe','ほのお'),
('1b72d661-b90f-4e67-853a-4cf428cc80fe','みず'),
('1b72d661-b90f-4e67-853a-4cf428cc80fe','かくとう'),
('1b72d661-b90f-4e67-853a-4cf428cc80fe','じめん'),
('efc65c97-d605-440b-a404-ba0ee0a647f1','ゴースト'),
('efc65c97-d605-440b-a404-ba0ee0a647f1','あく'),
('c81bab22-f182-4189-a735-8124d3f97e36','みず'),
('c81bab22-f182-4189-a735-8124d3f97e36','くさ'),
('c81bab22-f182-4189-a735-8124d3f97e36','かくとう'),
('c81bab22-f182-4189-a735-8124d3f97e36','じめん'),
('c81bab22-f182-4189-a735-8124d3f97e36','むし'),
('c81bab22-f182-4189-a735-8124d3f97e36','はがね'),
('c81bab22-f182-4189-a735-8124d3f97e36','フェアリー'),
('a399af77-e1c1-4dd9-a121-ac7e4eb88b88','どく'),
('a399af77-e1c1-4dd9-a121-ac7e4eb88b88','ひこう'),
('a399af77-e1c1-4dd9-a121-ac7e4eb88b88','むし'),
('8d6fc3aa-e3a9-4331-9351-c948131bc092','でんき'),
('8d6fc3aa-e3a9-4331-9351-c948131bc092','いわ'),
('5edae7d0-b84f-4df0-aa77-920bf1debf92','どく'),
('5edae7d0-b84f-4df0-aa77-920bf1debf92','ゴースト'),
('5edae7d0-b84f-4df0-aa77-920bf1debf92','はがね'),
('c148c1f5-9ce5-4d43-9680-4167c4fc3bf9','じめん'),
('efaf923e-b4a6-4608-9679-b7b619295020','みず'),
('efaf923e-b4a6-4608-9679-b7b619295020','じめん'),
('efaf923e-b4a6-4608-9679-b7b619295020','いわ'),
('00140872-80df-4bd8-bf9e-b182323ea665','こおり'),
('00140872-80df-4bd8-bf9e-b182323ea665','ドラゴン'),
('00140872-80df-4bd8-bf9e-b182323ea665','フェアリー'),
('3c6c90d7-ac08-4a20-9d7f-0fd1b2e5193d','くさ'),
('3c6c90d7-ac08-4a20-9d7f-0fd1b2e5193d','でんき'),
('3c6c90d7-ac08-4a20-9d7f-0fd1b2e5193d','かくとう'),
('3c6c90d7-ac08-4a20-9d7f-0fd1b2e5193d','むし'),
('3c6c90d7-ac08-4a20-9d7f-0fd1b2e5193d','フェアリー'),
('43a87790-13ce-4447-b29d-bcd654b902b0','くさ'),
('43a87790-13ce-4447-b29d-bcd654b902b0','でんき'),
('b1296a97-3798-45a2-99fb-d5a96bb21009','ほのお'),
('b1296a97-3798-45a2-99fb-d5a96bb21009','かくとう'),
('b1296a97-3798-45a2-99fb-d5a96bb21009','いわ'),
('b1296a97-3798-45a2-99fb-d5a96bb21009','はがね'),
('9c428543-a904-43f1-974d-59be53dd4ad6','ほのお'),
('9c428543-a904-43f1-974d-59be53dd4ad6','こおり'),
('9c428543-a904-43f1-974d-59be53dd4ad6','ひこう'),
('9c428543-a904-43f1-974d-59be53dd4ad6','エスパー'),
('2328ea34-e5a4-426a-a1d2-83240417f566','くさ'),
('ba937871-0fe6-4a60-a238-c86a016ed3a5','でんき'),
('ba937871-0fe6-4a60-a238-c86a016ed3a5','こおり'),
('ba937871-0fe6-4a60-a238-c86a016ed3a5','いわ'),
('ba937871-0fe6-4a60-a238-c86a016ed3a5','ゴースト'),
('ba937871-0fe6-4a60-a238-c86a016ed3a5','あく'),
('6a02e03e-30ca-464c-81d3-494e3cdb11e9','じめん'),
('57281b8f-330f-4d29-8047-905c67dc3194','ほのお'),
('57281b8f-330f-4d29-8047-905c67dc3194','じめん'),
('57281b8f-330f-4d29-8047-905c67dc3194','ゴースト'),
('57281b8f-330f-4d29-8047-905c67dc3194','あく'),
('2416f431-0e04-414f-bcd0-0b5002b47227','ほのお'),
('2416f431-0e04-414f-bcd0-0b5002b47227','かくとう'),
('2416f431-0e04-414f-bcd0-0b5002b47227','じめん'),
('cadb5595-86ba-4318-8b4c-73c124f32a13','みず'),
('cadb5595-86ba-4318-8b4c-73c124f32a13','くさ'),
('cadb5595-86ba-4318-8b4c-73c124f32a13','こおり'),
('4d2f2fd9-8725-4b3a-bbc5-719022061840','じめん'),
('efa33adb-af66-47c3-bce1-aa050212745a','じめん'),
('efa33adb-af66-47c3-bce1-aa050212745a','ひこう'),
('efa33adb-af66-47c3-bce1-aa050212745a','エスパー'),
('b6375a73-50ee-4469-ada1-d159d7053a10','ほのお'),
('b6375a73-50ee-4469-ada1-d159d7053a10','かくとう'),
('b6375a73-50ee-4469-ada1-d159d7053a10','どく'),
('b6375a73-50ee-4469-ada1-d159d7053a10','ひこう'),
('b6375a73-50ee-4469-ada1-d159d7053a10','むし'),
('b6375a73-50ee-4469-ada1-d159d7053a10','いわ'),
('b6375a73-50ee-4469-ada1-d159d7053a10','はがね'),
('4e894aee-54bb-45f3-912a-6be84e53d10a','ほのお'),
('4e894aee-54bb-45f3-912a-6be84e53d10a','かくとう'),
('4e894aee-54bb-45f3-912a-6be84e53d10a','むし'),
('4e894aee-54bb-45f3-912a-6be84e53d10a','いわ'),
('4e894aee-54bb-45f3-912a-6be84e53d10a','はがね'),
('4e894aee-54bb-45f3-912a-6be84e53d10a','フェアリー'),
('92c07146-a30b-4c59-82d4-df97e303ffee','みず'),
('92c07146-a30b-4c59-82d4-df97e303ffee','くさ'),
('92c07146-a30b-4c59-82d4-df97e303ffee','こおり'),
('92c07146-a30b-4c59-82d4-df97e303ffee','かくとう'),
('92c07146-a30b-4c59-82d4-df97e303ffee','じめん'),
('92c07146-a30b-4c59-82d4-df97e303ffee','はがね'),
('6630a4b7-0263-4551-a208-b0cd09204ec1','でんき'),
('6630a4b7-0263-4551-a208-b0cd09204ec1','こおり'),
('6630a4b7-0263-4551-a208-b0cd09204ec1','どく'),
('6630a4b7-0263-4551-a208-b0cd09204ec1','いわ'),
('6630a4b7-0263-4551-a208-b0cd09204ec1','はがね'),
('137ac9c8-ae5f-4787-8954-86be737ebba9','ほのお'),
('137ac9c8-ae5f-4787-8954-86be737ebba9','こおり'),
('137ac9c8-ae5f-4787-8954-86be737ebba9','どく'),
('137ac9c8-ae5f-4787-8954-86be737ebba9','ひこう'),
('137ac9c8-ae5f-4787-8954-86be737ebba9','むし'),
('d3ff6421-5d19-4e22-882f-459bc0fdec6b','ほのお'),
('d3ff6421-5d19-4e22-882f-459bc0fdec6b','かくとう'),
('d3ff6421-5d19-4e22-882f-459bc0fdec6b','いわ'),
('d3ff6421-5d19-4e22-882f-459bc0fdec6b','はがね'),
('ff0cb95d-e9f6-48a7-a0c2-19f688373f45','ほのお'),
('ff0cb95d-e9f6-48a7-a0c2-19f688373f45','みず'),
('ff0cb95d-e9f6-48a7-a0c2-19f688373f45','くさ'),
('ff0cb95d-e9f6-48a7-a0c2-19f688373f45','かくとう'),
('ff0cb95d-e9f6-48a7-a0c2-19f688373f45','はがね'),
('c1c8e0ed-e630-4a9d-9d4b-251c3e9400f2','ひこう'),
('c1c8e0ed-e630-4a9d-9d4b-251c3e9400f2','ゴースト'),
('c1c8e0ed-e630-4a9d-9d4b-251c3e9400f2','フェアリー'),
('bf7c3e54-bc09-4acf-8fd6-76a2fca0a150','ゴースト'),
('bf7c3e54-bc09-4acf-8fd6-76a2fca0a150','あく'),
('83c5f0d8-854a-4a78-af20-d3cffb5e2394','ほのお'),
('83c5f0d8-854a-4a78-af20-d3cffb5e2394','いわ'),
('83c5f0d8-854a-4a78-af20-d3cffb5e2394','ゴースト'),
('83c5f0d8-854a-4a78-af20-d3cffb5e2394','あく'),
('83c5f0d8-854a-4a78-af20-d3cffb5e2394','はがね'),
('e74cb2f1-e871-4b47-8169-816d1abcba97','じめん'),
('e74cb2f1-e871-4b47-8169-816d1abcba97','ゴースト'),
('e74cb2f1-e871-4b47-8169-816d1abcba97','あく'),
('d2f6e9b3-82c6-4f3c-b878-8a0805d08b94','みず'),
('d2f6e9b3-82c6-4f3c-b878-8a0805d08b94','じめん'),
('d2f6e9b3-82c6-4f3c-b878-8a0805d08b94','いわ'),
('0ac70058-b64d-4cf3-b30d-0da23dddb97d','くさ'),
('0ac70058-b64d-4cf3-b30d-0da23dddb97d','じめん'),
('13b2059a-9674-4d2a-92f9-d976118f164f','ほのお'),
('13b2059a-9674-4d2a-92f9-d976118f164f','かくとう'),
('13b2059a-9674-4d2a-92f9-d976118f164f','じめん'),
('13b2059a-9674-4d2a-92f9-d976118f164f','いわ'),
('62d2575c-d3f1-4221-b838-ed50f33fd97f','こおり'),
('62d2575c-d3f1-4221-b838-ed50f33fd97f','いわ'),
('07592851-333c-4851-9d65-7b18d875a130','ほのお'),
('07592851-333c-4851-9d65-7b18d875a130','こおり'),
('07592851-333c-4851-9d65-7b18d875a130','どく'),
('07592851-333c-4851-9d65-7b18d875a130','むし'),
('84f021c9-ab77-461b-80ae-96edf4fb4567','でんき'),
('84f021c9-ab77-461b-80ae-96edf4fb4567','こおり'),
('84f021c9-ab77-461b-80ae-96edf4fb4567','いわ'),
('84f021c9-ab77-461b-80ae-96edf4fb4567','ゴースト'),
('84f021c9-ab77-461b-80ae-96edf4fb4567','あく'),
('c2835a2a-a9e4-4b9b-9c6e-d1bc4979fcbf','ほのお'),
('c2835a2a-a9e4-4b9b-9c6e-d1bc4979fcbf','みず'),
('c2835a2a-a9e4-4b9b-9c6e-d1bc4979fcbf','かくとう'),
('c2835a2a-a9e4-4b9b-9c6e-d1bc4979fcbf','じめん'),
('fe62512b-3122-4cfb-b856-cfab6706ed06','ひこう'),
('fe62512b-3122-4cfb-b856-cfab6706ed06','エスパー'),
('fe62512b-3122-4cfb-b856-cfab6706ed06','フェアリー'),
('de35a6a7-ff48-48cb-82c1-210b16003b70','くさ'),
('13c246e3-1274-4674-ae2d-d68996181af5','ほのお'),
('13c246e3-1274-4674-ae2d-d68996181af5','こおり'),
('13c246e3-1274-4674-ae2d-d68996181af5','どく'),
('13c246e3-1274-4674-ae2d-d68996181af5','ひこう'),
('13c246e3-1274-4674-ae2d-d68996181af5','はがね'),
('1981290a-e4b6-4f7e-9ee3-1ced0a9ca359','ほのお'),
('1981290a-e4b6-4f7e-9ee3-1ced0a9ca359','かくとう'),
('1981290a-e4b6-4f7e-9ee3-1ced0a9ca359','いわ'),
('1981290a-e4b6-4f7e-9ee3-1ced0a9ca359','はがね'),
('5959d117-b13e-4a3c-bf5a-103b3da25d5d','かくとう'),
('5959d117-b13e-4a3c-bf5a-103b3da25d5d','ひこう'),
('5959d117-b13e-4a3c-bf5a-103b3da25d5d','フェアリー'),
('6bde063b-d78b-4791-a6ea-4f2c90de5059','でんき'),
('6bde063b-d78b-4791-a6ea-4f2c90de5059','こおり'),
('6bde063b-d78b-4791-a6ea-4f2c90de5059','いわ'),
('6bde063b-d78b-4791-a6ea-4f2c90de5059','ゴースト'),
('6bde063b-d78b-4791-a6ea-4f2c90de5059','あく'),
('04a7c71f-2e5e-4fb6-b249-6231a505a469','ゴースト'),
('04a7c71f-2e5e-4fb6-b249-6231a505a469','あく'),
('12101b9b-7e40-4298-a2ff-9d7ea2f60469','じめん'),
('12101b9b-7e40-4298-a2ff-9d7ea2f60469','エスパー'),
('32209607-554e-43d9-aca7-9e2a64cde7d2','むし'),
('32209607-554e-43d9-aca7-9e2a64cde7d2','ゴースト'),
('32209607-554e-43d9-aca7-9e2a64cde7d2','あく'),
('ba9f28e6-bea5-4b79-904e-165034f875d1','むし'),
('ba9f28e6-bea5-4b79-904e-165034f875d1','ゴースト'),
('ba9f28e6-bea5-4b79-904e-165034f875d1','あく'),
('bb5bbfed-3572-416b-aa66-1d0b2b14c45d','ほのお'),
('bb5bbfed-3572-416b-aa66-1d0b2b14c45d','かくとう'),
('bb5bbfed-3572-416b-aa66-1d0b2b14c45d','いわ'),
('bb5bbfed-3572-416b-aa66-1d0b2b14c45d','はがね'),
('b466ab80-8224-4e55-881a-708ed45989df','ほのお'),
('ded6f82a-ac2f-4cc5-8e4a-1ebd36ab8adb','くさ'),
('ded6f82a-ac2f-4cc5-8e4a-1ebd36ab8adb','でんき'),
('ded6f82a-ac2f-4cc5-8e4a-1ebd36ab8adb','ゴースト'),
('ded6f82a-ac2f-4cc5-8e4a-1ebd36ab8adb','あく'),
('b353c7f6-caf2-4215-a4ef-291348c32ce5','ほのお'),
('b353c7f6-caf2-4215-a4ef-291348c32ce5','いわ'),
('0c702f8c-7a84-4552-b21e-76434f0341ae','ほのお'),
('0c702f8c-7a84-4552-b21e-76434f0341ae','かくとう'),
('8b66568b-d635-4cb3-bc7b-51625a1dcb11','ほのお'),
('8b66568b-d635-4cb3-bc7b-51625a1dcb11','かくとう'),
('8b66568b-d635-4cb3-bc7b-51625a1dcb11','じめん'),
('0c9a22d2-ae05-40e7-8fc6-32fdb107afff','むし'),
('0c9a22d2-ae05-40e7-8fc6-32fdb107afff','ゴースト'),
('0c9a22d2-ae05-40e7-8fc6-32fdb107afff','あく'),
('bd5886db-ec28-482d-b621-95bb57f376bb','みず'),
('bd5886db-ec28-482d-b621-95bb57f376bb','じめん'),
('bd5886db-ec28-482d-b621-95bb57f376bb','いわ'),
('bd5886db-ec28-482d-b621-95bb57f376bb','ゴースト'),
('bd5886db-ec28-482d-b621-95bb57f376bb','あく'),
('bec88041-86b9-4a41-bff7-85c9fd2f7302','こおり'),
('bec88041-86b9-4a41-bff7-85c9fd2f7302','ドラゴン'),
('bec88041-86b9-4a41-bff7-85c9fd2f7302','フェアリー'),
('0764ccb7-750b-4c64-864d-ebd55a53cd5b','ほのお'),
('0764ccb7-750b-4c64-864d-ebd55a53cd5b','みず'),
('0764ccb7-750b-4c64-864d-ebd55a53cd5b','かくとう'),
('0764ccb7-750b-4c64-864d-ebd55a53cd5b','じめん'),
('07e63725-d7ec-47aa-85ce-993bd07c2d2c','みず'),
('07e63725-d7ec-47aa-85ce-993bd07c2d2c','くさ'),
('07e63725-d7ec-47aa-85ce-993bd07c2d2c','こおり'),
('07e63725-d7ec-47aa-85ce-993bd07c2d2c','ゴースト'),
('07e63725-d7ec-47aa-85ce-993bd07c2d2c','あく'),
('9b15b3e1-94b0-4a8f-85a3-78b50b290cad','ほのお'),
('9b15b3e1-94b0-4a8f-85a3-78b50b290cad','かくとう'),
('9b15b3e1-94b0-4a8f-85a3-78b50b290cad','じめん'),
('52a49909-7a60-4b96-ad56-1d87e775727b','でんき'),
('52a49909-7a60-4b96-ad56-1d87e775727b','こおり'),
('52a49909-7a60-4b96-ad56-1d87e775727b','いわ'),
('52a49909-7a60-4b96-ad56-1d87e775727b','フェアリー'),
('ea6a6007-40cc-47fc-ae09-63cb7b331363','ほのお'),
('77185f2d-86ed-454d-9f16-f7de2027dec9','こおり'),
('77185f2d-86ed-454d-9f16-f7de2027dec9','かくとう'),
('77185f2d-86ed-454d-9f16-f7de2027dec9','むし'),
('77185f2d-86ed-454d-9f16-f7de2027dec9','ドラゴン'),
('77185f2d-86ed-454d-9f16-f7de2027dec9','フェアリー'),
('605e5fed-2ec3-497f-a69a-4b5b05ad2373','みず'),
('605e5fed-2ec3-497f-a69a-4b5b05ad2373','くさ'),
('605e5fed-2ec3-497f-a69a-4b5b05ad2373','こおり'),
('605e5fed-2ec3-497f-a69a-4b5b05ad2373','かくとう'),
('082bc3a4-20a1-4dc7-b570-0aa307e88690','かくとう'),
('082bc3a4-20a1-4dc7-b570-0aa307e88690','ひこう'),
('082bc3a4-20a1-4dc7-b570-0aa307e88690','フェアリー'),
('3bd8e5e0-fed0-45d3-bd10-ee6f4faf8e62','ほのお'),
('3bd8e5e0-fed0-45d3-bd10-ee6f4faf8e62','じめん'),
('3bd8e5e0-fed0-45d3-bd10-ee6f4faf8e62','ゴースト'),
('3bd8e5e0-fed0-45d3-bd10-ee6f4faf8e62','あく'),
('ec887d64-6852-49b4-a0ed-885101b44dc2','むし'),
('ec887d64-6852-49b4-a0ed-885101b44dc2','フェアリー'),
('506c86ee-fe70-4856-bb34-166b135a5e9a','くさ'),
('506c86ee-fe70-4856-bb34-166b135a5e9a','でんき'),
('506c86ee-fe70-4856-bb34-166b135a5e9a','かくとう'),
('506c86ee-fe70-4856-bb34-166b135a5e9a','じめん'),
('9a16ef20-2c55-4d99-b8f8-3ee76b85d171','かくとう'),
('9a16ef20-2c55-4d99-b8f8-3ee76b85d171','じめん'),
('025c3b5e-30a5-490d-8b77-24a5813caaf9','どく'),
('025c3b5e-30a5-490d-8b77-24a5813caaf9','はがね'),
('ab2ffa3b-6a58-491b-a1c8-8011fdab9f8c','でんき'),
('ab2ffa3b-6a58-491b-a1c8-8011fdab9f8c','こおり'),
('ab2ffa3b-6a58-491b-a1c8-8011fdab9f8c','ひこう'),
('ab2ffa3b-6a58-491b-a1c8-8011fdab9f8c','エスパー'),
('ab2ffa3b-6a58-491b-a1c8-8011fdab9f8c','フェアリー'),
('99732dde-ca39-4b07-8526-1b90ed24fb8b','こおり'),
('99732dde-ca39-4b07-8526-1b90ed24fb8b','ドラゴン'),
('99732dde-ca39-4b07-8526-1b90ed24fb8b','フェアリー'),
('2bcb75aa-4089-428c-8dd4-5b0864d7ffac','ほのお'),
('2bcb75aa-4089-428c-8dd4-5b0864d7ffac','こおり'),
('2bcb75aa-4089-428c-8dd4-5b0864d7ffac','ひこう'),
('2bcb75aa-4089-428c-8dd4-5b0864d7ffac','ゴースト'),
('2bcb75aa-4089-428c-8dd4-5b0864d7ffac','あく'),
('6ef79213-8d09-48a9-bd19-48e289193af9','ほのお'),
('6ef79213-8d09-48a9-bd19-48e289193af9','こおり'),
('6ef79213-8d09-48a9-bd19-48e289193af9','ひこう'),
('6ef79213-8d09-48a9-bd19-48e289193af9','ゴースト'),
('6ef79213-8d09-48a9-bd19-48e289193af9','あく'),
('8f34b13b-99a3-4308-b493-6def4204afbd','こおり'),
('8f34b13b-99a3-4308-b493-6def4204afbd','いわ'),
('8f34b13b-99a3-4308-b493-6def4204afbd','ドラゴン'),
('8f34b13b-99a3-4308-b493-6def4204afbd','フェアリー'),
('f7aebf3f-4272-4b6b-83fc-e950b2606413','ほのお'),
('f7aebf3f-4272-4b6b-83fc-e950b2606413','いわ'),
('e482dd5e-a905-4ae6-a97d-b8ab4827a748','でんき'),
('e482dd5e-a905-4ae6-a97d-b8ab4827a748','じめん'),
('e482dd5e-a905-4ae6-a97d-b8ab4827a748','エスパー'),
('98c91a82-e550-4f13-aee1-4da2d14ce87f','でんき'),
('98c91a82-e550-4f13-aee1-4da2d14ce87f','ひこう'),
('98c91a82-e550-4f13-aee1-4da2d14ce87f','いわ'),
('ab1fe650-0aa7-46f9-973b-e0e8fb1ed11a','みず'),
('ab1fe650-0aa7-46f9-973b-e0e8fb1ed11a','じめん'),
('ab1fe650-0aa7-46f9-973b-e0e8fb1ed11a','エスパー'),
('ab1fe650-0aa7-46f9-973b-e0e8fb1ed11a','いわ'),
('f5d64c0c-5e59-4d74-aba8-9ebb0aafb6d7','かくとう'),
('f5d64c0c-5e59-4d74-aba8-9ebb0aafb6d7','ひこう'),
('f5d64c0c-5e59-4d74-aba8-9ebb0aafb6d7','エスパー'),
('f5d64c0c-5e59-4d74-aba8-9ebb0aafb6d7','フェアリー'),
('31027691-7f50-4084-9107-d03ff1f6faa2','でんき'),
('31027691-7f50-4084-9107-d03ff1f6faa2','ひこう'),
('31027691-7f50-4084-9107-d03ff1f6faa2','いわ'),
('462331c1-afab-4d8a-ac53-ebdcfc00a75a','くさ'),
('462331c1-afab-4d8a-ac53-ebdcfc00a75a','でんき'),
('f8f211f0-81f6-417f-866f-6142b1fee99b','ゴースト'),
('f8f211f0-81f6-417f-866f-6142b1fee99b','はがね'),
('04348fa1-4b27-47d8-9dbd-42b222b5ad79','こおり'),
('04348fa1-4b27-47d8-9dbd-42b222b5ad79','ひこう'),
('04348fa1-4b27-47d8-9dbd-42b222b5ad79','エスパー'),
('04348fa1-4b27-47d8-9dbd-42b222b5ad79','ドラゴン'),
('04348fa1-4b27-47d8-9dbd-42b222b5ad79','フェアリー'),
('ab536bf7-c109-4577-ac25-294715df185d','ほのお'),
('ab536bf7-c109-4577-ac25-294715df185d','こおり'),
('ab536bf7-c109-4577-ac25-294715df185d','どく'),
('ab536bf7-c109-4577-ac25-294715df185d','ひこう'),
('ab536bf7-c109-4577-ac25-294715df185d','むし'),
('ca023531-0121-4cf3-8345-3f1d01af335b','みず'),
('ca023531-0121-4cf3-8345-3f1d01af335b','じめん'),
('ca023531-0121-4cf3-8345-3f1d01af335b','いわ'),
('13bbae82-6acf-41b6-9a57-7e1dba8f1741','くさ'),
('13bbae82-6acf-41b6-9a57-7e1dba8f1741','でんき'),
('c0655ae7-0cb9-4388-9ae9-e910d34a9faf','かくとう'),
('fe75cf44-a6d5-4771-9544-401da2fa18d3','ほのお'),
('fe75cf44-a6d5-4771-9544-401da2fa18d3','でんき'),
('4a7e6d50-8185-4279-95cd-bc0669399be8','くさ'),
('4a7e6d50-8185-4279-95cd-bc0669399be8','でんき'),
('4a7e6d50-8185-4279-95cd-bc0669399be8','かくとう'),
('4a7e6d50-8185-4279-95cd-bc0669399be8','じめん'),
('e8578ada-4955-4a00-b7a8-a942008b075b','じめん'),
('f44c378a-4672-461f-bdb8-297037de6d34','みず'),
('f44c378a-4672-461f-bdb8-297037de6d34','かくとう'),
('f44c378a-4672-461f-bdb8-297037de6d34','じめん'),
('f44c378a-4672-461f-bdb8-297037de6d34','いわ'),
('f5c1b645-ae5a-49a6-a886-69dd13fce9ac','こおり'),
('f5c1b645-ae5a-49a6-a886-69dd13fce9ac','どく'),
('f5c1b645-ae5a-49a6-a886-69dd13fce9ac','ひこう'),
('f5c1b645-ae5a-49a6-a886-69dd13fce9ac','むし'),
('f5c1b645-ae5a-49a6-a886-69dd13fce9ac','ドラゴン'),
('f5c1b645-ae5a-49a6-a886-69dd13fce9ac','フェアリー'),
('ef7631a6-005e-4b32-92b0-5c209d5f9ee3','こおり'),
('ef7631a6-005e-4b32-92b0-5c209d5f9ee3','どく'),
('ef7631a6-005e-4b32-92b0-5c209d5f9ee3','ひこう'),
('ef7631a6-005e-4b32-92b0-5c209d5f9ee3','むし'),
('ef7631a6-005e-4b32-92b0-5c209d5f9ee3','ドラゴン'),
('ef7631a6-005e-4b32-92b0-5c209d5f9ee3','フェアリー'),
('59f83115-3405-4cfc-bdb5-67973bc91b6e','じめん'),
('59f83115-3405-4cfc-bdb5-67973bc91b6e','エスパー'),
('8f3487f7-f245-47cc-b363-76ea13dbd4b4','みず'),
('8f3487f7-f245-47cc-b363-76ea13dbd4b4','ひこう'),
('8f3487f7-f245-47cc-b363-76ea13dbd4b4','いわ'),
('5a783b23-fdce-4ac8-9834-bb1ce241e2ae','ゴースト'),
('5a783b23-fdce-4ac8-9834-bb1ce241e2ae','あく'),
('310a2020-bd16-4857-8fa5-f18e0b5e8086','どく'),
('310a2020-bd16-4857-8fa5-f18e0b5e8086','ゴースト'),
('310a2020-bd16-4857-8fa5-f18e0b5e8086','はがね'),
('3992df26-179d-4f6d-b3c7-01e00b3f3212','どく'),
('3992df26-179d-4f6d-b3c7-01e00b3f3212','はがね'),
('3992df26-179d-4f6d-b3c7-01e00b3f3212','フェアリー'),
('6c3a6316-f7e8-4e3e-b676-dde463399d59','かくとう'),
('6c3a6316-f7e8-4e3e-b676-dde463399d59','むし'),
('6c3a6316-f7e8-4e3e-b676-dde463399d59','フェアリー'),
('a5945afe-be2b-4e46-bce7-3da763819924','ほのお'),
('a5945afe-be2b-4e46-bce7-3da763819924','かくとう'),
('a5945afe-be2b-4e46-bce7-3da763819924','じめん'),
('198f08d7-6ec0-4986-9634-b87283ae6b3a','ゴースト'),
('198f08d7-6ec0-4986-9634-b87283ae6b3a','あく'),
('4ce0b088-b80c-46aa-b64c-807c885de3db','ひこう'),
('4ce0b088-b80c-46aa-b64c-807c885de3db','エスパー'),
('4ce0b088-b80c-46aa-b64c-807c885de3db','フェアリー'),
('fe600a0a-1488-4242-ad54-1fada72eca37','ほのお'),
('fe600a0a-1488-4242-ad54-1fada72eca37','むし'),
('fe600a0a-1488-4242-ad54-1fada72eca37','いわ'),
('fe600a0a-1488-4242-ad54-1fada72eca37','ゴースト'),
('fe600a0a-1488-4242-ad54-1fada72eca37','あく'),
('fe600a0a-1488-4242-ad54-1fada72eca37','はがね'),
('630545ac-b378-4948-a2cc-232b87987ca9','みず'),
('630545ac-b378-4948-a2cc-232b87987ca9','くさ'),
('630545ac-b378-4948-a2cc-232b87987ca9','こおり'),
('630545ac-b378-4948-a2cc-232b87987ca9','ゴースト'),
('630545ac-b378-4948-a2cc-232b87987ca9','あく'),
('fc29bbab-10d7-47e4-9ffd-a7b4a953cb3c','ほのお'),
('fc29bbab-10d7-47e4-9ffd-a7b4a953cb3c','ひこう'),
('fc29bbab-10d7-47e4-9ffd-a7b4a953cb3c','いわ'),
('fc29bbab-10d7-47e4-9ffd-a7b4a953cb3c','はがね'),
('ccce3d41-3658-4956-bd57-11d304fb32c6','ほのお'),
('ccce3d41-3658-4956-bd57-11d304fb32c6','かくとう'),
('ccce3d41-3658-4956-bd57-11d304fb32c6','いわ'),
('ccce3d41-3658-4956-bd57-11d304fb32c6','はがね'),
('f1a6995a-7088-4571-8f73-b025694cc02b','かくとう'),
('f1a6995a-7088-4571-8f73-b025694cc02b','むし'),
('f1a6995a-7088-4571-8f73-b025694cc02b','フェアリー'),
('cedda196-f965-4765-a6a3-98f1c678ead3','ほのお'),
('cedda196-f965-4765-a6a3-98f1c678ead3','かくとう'),
('cedda196-f965-4765-a6a3-98f1c678ead3','じめん'),
('ee46d662-3075-4be5-97af-46777cce91ea','こおり'),
('ee46d662-3075-4be5-97af-46777cce91ea','じめん'),
('ee46d662-3075-4be5-97af-46777cce91ea','ドラゴン'),
('ee46d662-3075-4be5-97af-46777cce91ea','フェアリー'),
('26192261-290f-4337-bf17-5686341e3a85','ほのお'),
('26192261-290f-4337-bf17-5686341e3a85','かくとう'),
('26192261-290f-4337-bf17-5686341e3a85','じめん'),
('26192261-290f-4337-bf17-5686341e3a85','いわ'),
('531205fd-9bdc-4885-97e2-441a87e1bcb1','ドラゴン'),
('531205fd-9bdc-4885-97e2-441a87e1bcb1','フェアリー'),
('3da3ef08-c411-477b-9dfd-956d9ee7b1c5','くさ'),
('3da3ef08-c411-477b-9dfd-956d9ee7b1c5','でんき'),
('3da3ef08-c411-477b-9dfd-956d9ee7b1c5','かくとう'),
('3da3ef08-c411-477b-9dfd-956d9ee7b1c5','いわ'),
('cde790f8-07fe-4ef2-8568-1241f05b9919','かくとう'),
('cde790f8-07fe-4ef2-8568-1241f05b9919','じめん'),
('fa85a98d-a7fd-4600-aff4-41cce9588f56','こおり'),
('fa85a98d-a7fd-4600-aff4-41cce9588f56','ゴースト'),
('fa85a98d-a7fd-4600-aff4-41cce9588f56','ドラゴン'),
('fa85a98d-a7fd-4600-aff4-41cce9588f56','あく'),
('fa85a98d-a7fd-4600-aff4-41cce9588f56','フェアリー'),
('d61dc2bf-1519-4d78-8325-92170d5f2900','かくとう'),
('d61dc2bf-1519-4d78-8325-92170d5f2900','みず'),
('d61dc2bf-1519-4d78-8325-92170d5f2900','くさ'),
('d61dc2bf-1519-4d78-8325-92170d5f2900','はがね'),
('d61dc2bf-1519-4d78-8325-92170d5f2900','じめん');
`
