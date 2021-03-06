package samples

// SingleFish returns a simple fish sample json string
func SingleFish() string {
	return `{
		"id":1,
		"file-name":"bitterling",
		"name":{
			"name-USen": "bitterling", 
			"name-EUen":"bitterling",
			"name-EUde":"Bitterling",
			"name-EUes":"amarguillo",
			"name-USes":"amarguillo",
			"name-EUfr":"bouvière",
			"name-USfr":"bouvière",
			"name-EUit":"rodeo",
			"name-EUnl":"bittervoorn",
			"name-CNzh":"红目鲫",
			"name-TWzh":"紅目鯽",
			"name-JPja":"タナゴ",
			"name-KRko":"납줄개",
			"name-EUru":"горчак"
		},
		"availability":{
			"month-northern":"11-3",
			"month-southern":"5-9",
			"time":"",
			"isAllDay":true,
			"isAllYear":false,
			"location":"River",
			"rarity":"Common",
			"month-array-northern":[11,12,1,2,3],
			"month-array-southern":[5,6,7,8,9],
			"time-array":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]
		},
		"shadow":"Smallest (1)",
		"price":900,
		"price-cj":1350,
		"catch-phrase":"I caught a bitterling! It's mad at me, but only a little.",
		"museum-phrase":"Bitterlings hide their eggs inside large bivalves—like clams—where the young can stay safe until grown. The bitterling isn't being sneaky. No, their young help keep the bivalve healthy by eating invading parasites! It's a wonderful bit of evolutionary deal making, don't you think? Each one keeping the other safe... Though eating parasites does not sound like a happy childhood... Is that why the fish is so bitter?",
		"image_uri":"https://acnhapi.com/v1/images/fish/1",
		"icon_uri":"https://acnhapi.com/v1/icons/fish/1"
		}`
}

// MultipleFish returns multiple (2) simple fish
func MultipleFish() string {
	return `
	[
		{
			"id":1,
			"file-name":"bitterling",
			"name":{
				"name-USen": "bitterling", 
				"name-EUen":"bitterling",
				"name-EUde":"Bitterling",
				"name-EUes":"amarguillo",
				"name-USes":"amarguillo",
				"name-EUfr":"bouvière",
				"name-USfr":"bouvière",
				"name-EUit":"rodeo",
				"name-EUnl":"bittervoorn",
				"name-CNzh":"红目鲫",
				"name-TWzh":"紅目鯽",
				"name-JPja":"タナゴ",
				"name-KRko":"납줄개",
				"name-EUru":"горчак"
			},
			"availability":{
				"month-northern":"11-3",
				"month-southern":"5-9",
				"time":"",
				"isAllDay":true,
				"isAllYear":false,
				"location":"River",
				"rarity":"Common",
				"month-array-northern":[11,12,1,2,3],
				"month-array-southern":[5,6,7,8,9],
				"time-array":[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]
			},
			"shadow":"Smallest (1)",
			"price":900,
			"price-cj":1350,
			"catch-phrase":"I caught a bitterling! It's mad at me, but only a little.",
			"museum-phrase":"Bitterlings hide their eggs inside large bivalves—like clams—where the young can stay safe until grown. The bitterling isn't being sneaky. No, their young help keep the bivalve healthy by eating invading parasites! It's a wonderful bit of evolutionary deal making, don't you think? Each one keeping the other safe... Though eating parasites does not sound like a happy childhood... Is that why the fish is so bitter?",
			"image_uri":"https://acnhapi.com/v1/images/fish/1",
			"icon_uri":"https://acnhapi.com/v1/icons/fish/1"
		},

		{
			"id": 2,
			"file-name": "pale_chub",
			"name": {
			  "name-USen": "pale chub",
			  "name-EUen": "pale chub",
			  "name-EUde": "Döbel",
			  "name-EUes": "cacho",
			  "name-USes": "cacho",
			  "name-EUfr": "chevaine",
			  "name-USfr": "chevaine",
			  "name-EUit": "zacco",
			  "name-EUnl": "Japanse drakenvis",
			  "name-CNzh": "溪哥",
			  "name-TWzh": "溪哥",
			  "name-JPja": "オイカワ",
			  "name-KRko": "피라미",
			  "name-EUru": "светлый закко"
			},
			"availability": {
			  "month-northern": "",
			  "month-southern": "",
			  "time": "9am - 4pm",
			  "isAllDay": false,
			  "isAllYear": true,
			  "location": "River",
			  "rarity": "Common",
			  "month-array-northern": [1,2,3,4],
			  "month-array-southern": [9,10,11,12],
			  "time-array": [9,10,11,12,13,14,15]
			},
			"shadow": "Smallest (1)",
			"price": 200,
			"price-cj": 300,
			"catch-phrase": "I caught a pale chub! That name seems a bit judgy...",
			"museum-phrase": "The pale chub is a river fish with simple back-and-silver coloring. Interestingly, the males' coloring becomes most vibrant when he's trying to attract a mate! Though normally plain, these fellows really know how to look stylish when they want to. Perhaps I could learn a thing or two from the pale chub...",
			"image_uri": "https://acnhapi.com/v1/images/fish/2",
			"icon_uri": "https://acnhapi.com/v1/icons/fish/2"
		  }
	]`
}

// ComplexItems returns multiple items (fish and bugs) in a json
// string that contain edge conditiions, incorrect data (known and
// fixed in conversion) and accurate info
func ComplexItems() string {
	return `
	[
		{
			"id": 17,
			"file-name": "loach",
			"name": {
				"name-USen": "loach",
				"name-EUen": "loach",
				"name-EUde": "Bachschmerle",
				"name-EUes": "locha",
				"name-USes": "locha",
				"name-EUfr": "loche d'étang",
				"name-USfr": "loche d'étang",
				"name-EUit": "cobite",
				"name-EUnl": "modderkruiper",
				"name-CNzh": "泥鳅",
				"name-TWzh": "泥鰍",
				"name-JPja": "ドジョウ",
				"name-KRko": "미꾸라지",
				"name-EUru": "вьюн"
			},
			"availability": {
				"month-northern": "3-5",
				"month-southern": "9-11",
				"time": "",
				"isAllDay": true,
				"isAllYear": false,
				"location": "River",
				"rarity": "Common",
				"month-array-northern": [
					3,
					4,
					5
				],
				"month-array-southern": [
					9,
					10,
					11
				],
				"time-array": [
					0,
					1,
					2,
					3,
					4,
					5,
					6,
					7,
					8,
					9,
					10,
					11,
					12,
					13,
					14,
					15,
					16,
					17,
					18,
					19,
					20,
					21,
					22,
					23
				]
			},
			"shadow": "Small (2)",
			"price": 400,
			"price-cj": 600,
			"catch-phrase": "I caught a loach! It's...looking at me with reproach.",
			"museum-phrase": "Loaches are a large group of bottom-feeding freshwater fish. Because there are so many species, it is difficult to generalize about them all. The one thing they all have in common is perhaps their most unappealing trait; the name 'loach.' Repellent, isn't it? They should have just called it the 'cockloach' and been done with it, eh wot! Loach...loach... It just tastes terrible in the beak.",
			"image_uri": "https://acnhapi.com/v1/images/fish/17",
			"icon_uri": "https://acnhapi.com/v1/icons/fish/17"
		},
		{
			"id": 43,
			"file-name": "ladybug",
			"name": {
				"name-USen": "ladybug",
				"name-EUen": "ladybug",
				"name-EUde": "Marienkäfer",
				"name-EUes": "mariquita",
				"name-USes": "catarina",
				"name-EUfr": "coccinelle",
				"name-USfr": "coccinelle",
				"name-EUit": "coccinella",
				"name-EUnl": "lieveheersbeestje",
				"name-CNzh": "瓢虫",
				"name-TWzh": "瓢蟲",
				"name-JPja": "テントウムシ",
				"name-KRko": "무당벌레",
				"name-EUru": "божья коровка"
			},
			"availability": {
				"month-northern": "3-6 & 10",
				"month-southern": "9-12 & 4",
				"time": "8am - 5pm",
				"isAllDay": false,
				"isAllYear": false,
				"location": "On flowers",
				"rarity": "Common",
				"month-array-northern": [
					3,
					4,
					5,
					6,
					10
				],
				"month-array-southern": [
					4,
					9,
					10,
					11,
					12
				],
				"time-array": [
					8,
					9,
					10,
					11,
					12,
					13,
					14,
					15,
					16
				]
			},
			"price": 200,
			"price-flick": 300,
			"catch-phrase": "I caught a ladybug! Sorry to disturb you, ma'am.",
			"museum-phrase": "Yes. Yes. Ladybugs are quite beloved thanks to their tiny round shape and adorable spots. To that I say BAH! The fact of the matter is, some have stripes instead of spots. And SOME have no markings at all! No, I shall never understand why it is said that when a ladybug lands on you...you'll have good luck. I shall only have a fainting spell.",
			"image_uri": "https://acnhapi.com/v1/images/bugs/43",
			"icon_uri": "https://acnhapi.com/v1/icons/bugs/43"
		},
		{
			"id": 17,
			"file-name": "abalone",
			"name": {
				"name-USen": "abalone",
				"name-EUen": "abalone",
				"name-EUde": "Seeohr",
				"name-EUnl": "abalone",
				"name-EUes": "abulón",
				"name-USes": "abulón",
				"name-EUfr": "ormeau",
				"name-USfr": "ormeau",
				"name-EUit": "abalone",
				"name-CNzh": "鲍鱼",
				"name-TWzh": "鮑魚",
				"name-JPja": "アワビ",
				"name-KRko": "전복",
				"name-EUru": "абалон"
			},
			"availability": {
				"month-northern": "6-1",
				"month-southern": "12-7",
				"time": "4pm - 9am",
				"isAllDay": false,
				"isAllYear": false,
				"month-array-northern": [
					6,
					7,
					8,
					9,
					10,
					11,
					12,
					1
				],
				"month-array-southern": [
					12,
					1,
					2,
					7,
					4,
					5,
					6,
					7
				],
				"time-array": [
					16,
					17,
					18,
					19,
					20,
					21,
					22,
					23,
					0,
					1,
					2,
					3,
					4,
					5,
					6,
					7,
					8
				]
			},
			"speed": "Medium",
			"shadow": "Medium",
			"price": 2000,
			"catch-phrase": "I got an abalone! Why do I want a sandwich now?",
			"image_uri": "https://acnhapi.com/v1/images/sea/17",
			"icon_uri": "https://acnhapi.com/v1/icons/sea/17",
			"museum-phrase": "My feathers, but the abalone is a most deceptive sea snail, wot wot! After seeing its dull brown shell from the outside, one would think it quite plain.But take a gander inside, and you'll see that the abalone's home is a shimmering beauty to behold! The inner layer of the shell is made of \"nacre,\" or \"mother-of-pearl,\" and, hoo, what a dazzling iridescent hue! Let this be a lesson! You mustn't judge a sea snail by the outside of its shell. Judge it by the inside of its shell instead."
		},
		{
			"id": 59,
			"file-name": "sea_bass",
			"name": {
				"name-USen": "sea bass",
				"name-EUen": "sea bass",
				"name-EUde": "Seebarsch",
				"name-EUes": "lubina",
				"name-USes": "lubina",
				"name-EUfr": "bar commun",
				"name-USfr": "bar commun",
				"name-EUit": "spigola",
				"name-EUnl": "zeebaars",
				"name-CNzh": "鲈鱼",
				"name-TWzh": "鱸魚",
				"name-JPja": "スズキ",
				"name-KRko": "농어",
				"name-EUru": "морской судак"
			},
			"availability": {
				"month-northern": "",
				"month-southern": "",
				"time": "",
				"isAllDay": true,
				"isAllYear": true,
				"location": "Sea",
				"rarity": "Common",
				"month-array-northern": [
					1,
					2,
					3,
					4,
					5,
					6,
					7,
					8,
					9,
					10,
					11,
					12
				],
				"month-array-southern": [
					1,
					2,
					3,
					4,
					5,
					6,
					7,
					8,
					9,
					10,
					11,
					12
				],
				"time-array": [
					0,
					1,
					2,
					3,
					4,
					5,
					6,
					7,
					8,
					9,
					10,
					11,
					12,
					13,
					14,
					15,
					16,
					17,
					18,
					19,
					20,
					21,
					22,
					23
				]
			},
			"shadow": "Large (5)",
			"price": 400,
			"price-cj": 600,
			"catch-phrase": "I caught a sea bass! No, wait- it's at least a C+!",
			"museum-phrase": "Sea bass is a name given to a variety of different species of saltwater fish. They are a varied bunch with some as small as four inches and some as staggeringly ginormous as eight feet! 'Sea bass' is a bit pedestrian though. Many species have better names, such as 'redbanded perch.' Or the delightfully whimsical dusky grouper! Or the potato cod! WILL NO ONE THINK OF THE PINK MAOMAO?!",
			"image_uri": "https://acnhapi.com/v1/images/fish/59",
			"icon_uri": "https://acnhapi.com/v1/icons/fish/59"
		},
		{
			"id": 24,
			"file-name": "honeybee",
			"name": {
				"name-USen": "honeybee",
				"name-EUen": "honeybee",
				"name-EUde": "Honigbiene",
				"name-EUes": "abeja melífera",
				"name-USes": "abeja melífera",
				"name-EUfr": "abeille naine",
				"name-USfr": "abeille naine",
				"name-EUit": "ape operaia",
				"name-EUnl": "honingbij",
				"name-CNzh": "蜜蜂",
				"name-TWzh": "蜜蜂",
				"name-JPja": "ミツバチ",
				"name-KRko": "꿀벌",
				"name-EUru": "восковая пчела"
			},
			"availability": {
				"month-northern": "3-7",
				"month-southern": "9-1",
				"time": "8am - 5pm",
				"isAllDay": false,
				"isAllYear": false,
				"location": "Flying",
				"rarity": "Common",
				"month-array-northern": [
					3,
					4,
					5,
					6,
					7
				],
				"month-array-southern": [
					9,
					10,
					11,
					12,
					1
				],
				"time-array": [
					8,
					9,
					10,
					11,
					12,
					13,
					14,
					15,
					16
				]
			},
			"price": 200,
			"price-flick": 300,
			"catch-phrase": "I caught a honeybee! Ah, sweet success!",
			"museum-phrase": "Did you know it takes a team of honeybees working together to transform flower nectar into honey? Indeed, forager bees suck nectar from flowers into their \"honey stomachs\" and then fly it to the hive. Hive bees then chew the substance and spit it into the honeycomb, fluttering their wings to dry it out. Yes, you could say honey is a tasty tribute to the hard work of the humble honeybee. Oh! Oh my! You mustn't confuse my lengthy description for admiration! At the end of the day, honeybees are still insects, and thus still ghastly! A wee bit less ghastly than most, I admit.",
			"image_uri": "https://acnhapi.com/v1/images/bugs/24",
			"icon_uri": "https://acnhapi.com/v1/icons/bugs/24"
		},
		{
			"id": 70,
			"file-name": "ocean_sunfish",
			"name": {
				"name-USen": "ocean sunfish",
				"name-EUen": "ocean sunfish",
				"name-EUde": "Mondfisch",
				"name-EUes": "pez luna",
				"name-USes": "pez luna",
				"name-EUfr": "lune de mer",
				"name-USfr": "poisson-lune",
				"name-EUit": "pesce luna",
				"name-EUnl": "klompvis",
				"name-CNzh": "翻车鱼",
				"name-TWzh": "翻車魚",
				"name-JPja": "マンボウ",
				"name-KRko": "개복치",
				"name-EUru": "луна-рыба"
			},
			"availability": {
				"month-northern": "7-9",
				"month-southern": "1-3",
				"time": "4am - 9pm",
				"isAllDay": false,
				"isAllYear": false,
				"location": "Sea",
				"rarity": "Uncommon",
				"month-array-northern": [
					7,
					8,
					9
				],
				"month-array-southern": [
					1,
					2,
					3
				],
				"time-array": [
					4,
					5,
					6,
					7,
					8,
					9,
					10,
					11,
					12,
					13,
					14,
					15,
					16,
					17,
					18,
					19,
					20
				]
			},
			"shadow": "Largest with fin (6)",
			"price": 4000,
			"price-cj": 6000,
			"catch-phrase": "I caught an ocean sunfish! Good thing I'm wearing ocean sunscreen!",
			"museum-phrase": "The ocean sunfish is a large relative of the blowfish with an unusual shape, like a fish head with a tail. They are a fairly relaxed species, often content to ride where the currents take them. This is fortunate as, otherwise, the sight of a large head coming toward you might be alarming!",
			"image_uri": "https://acnhapi.com/v1/images/fish/70",
			"icon_uri": "https://acnhapi.com/v1/icons/fish/70"
		}
	]`
}
