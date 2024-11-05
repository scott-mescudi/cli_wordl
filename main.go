package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	Red    = "\033[1;31m"
	Green  = "\033[1;32m"
	Magenta = "\033[1;35m"
	Reset  = "\033[0m"

	tr = "┐"
	tl = "┌"

	br = "┘"
	bl = "└"
)

func clearTerminal() {
	var cmd *exec.Cmd

	
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") 
	default:
		cmd = exec.Command("clear") 
	}

	cmd.Stdout = os.Stdout 
	cmd.Run()
}


func getRandomWord() (string, error) {
	file, err := os.Open("wordlists/5_letter_words.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	x := rand.Intn(3088)

	var (
		count int
		word  string
	)

	for scanner.Scan() {
		if count == x {
			word = scanner.Text()
			break
		}
		count++
	}

	return word, err
}

func help() {
	box := "██"
	fmt.Printf("%v%v%v = letter is in right spot\n%v%v%v = letter is in word but not in right spot\n%v%v%v = letter is not in word\ntype ':q' to exit\n\n", Green, box, Reset, Magenta, box, Reset, Red, box, Reset)

}

var words = []string{"aargh",
"aback",
"abbey",
"abbot",
"abide",
"abode",
"abort",
"about",
"above",
"abuse",
"ached",
"aches",
"acids",
"acing",
"acorn",
"acres",
"acted",
"actin",
"actor",
"acute",
"adage",
"adapt",
"added",
"adept",
"admit",
"adobe",
"adopt",
"adore",
"adult",
"afoot",
"after",
"again",
"agent",
"agile",
"aging",
"agony",
"agree",
"ahead",
"ahold",
"aided",
"aides",
"aimed",
"aisle",
"alamo",
"alarm",
"album",
"alert",
"algae",
"alias",
"alibi",
"alien",
"alike",
"alive",
"allee",
"alley",
"allow",
"aloft",
"aloha",
"alone",
"along",
"aloud",
"alpha",
"altar",
"alter",
"amaze",
"amber",
"amend",
"amigo",
"amino",
"amiss",
"among",
"ample",
"amply",
"amuse",
"angel",
"anger",
"angle",
"angry",
"angst",
"anise",
"ankle",
"annex",
"annoy",
"annul",
"antsy",
"anvil",
"apart",
"apple",
"apply",
"apron",
"aptly",
"arbor",
"areas",
"arena",
"argon",
"argue",
"ariel",
"arise",
"armed",
"armor",
"aroma",
"arose",
"array",
"arrow",
"arson",
"artsy",
"ascot",
"ashes",
"aside",
"asked",
"askew",
"aspen",
"asses",
"asset",
"atlas",
"attic",
"audio",
"audit",
"auger",
"aught",
"aunts",
"auras",
"avoid",
"await",
"awake",
"award",
"aware",
"awful",
"awoke",
"babes",
"backs",
"bacon",
"badge",
"badly",
"bagel",
"baggy",
"bails",
"baked",
"baker",
"bakes",
"balls",
"banal",
"bands",
"bangs",
"banjo",
"banks",
"barbs",
"bared",
"barge",
"barks",
"baron",
"barre",
"based",
"bases",
"basic",
"basil",
"basin",
"basis",
"baste",
"batch",
"bates",
"bathe",
"baths",
"baton",
"batty",
"bawdy",
"bayou",
"beach",
"beads",
"beams",
"beans",
"beard",
"bears",
"beast",
"beats",
"becks",
"beech",
"beefs",
"beefy",
"beeps",
"beers",
"beery",
"beets",
"began",
"begat",
"begin",
"begun",
"beige",
"being",
"belie",
"belle",
"bells",
"belly",
"below",
"belts",
"bench",
"bends",
"bendy",
"benes",
"benny",
"beret",
"berry",
"bialy",
"bible",
"bigot",
"bijou",
"biker",
"bikes",
"bilge",
"bills",
"billy",
"bimbo",
"binds",
"binge",
"bingo",
"birch",
"birds",
"birth",
"bison",
"bitch",
"bites",
"bitsy",
"bitty",
"black",
"blade",
"blame",
"bland",
"blank",
"blast",
"blaze",
"bleak",
"bleed",
"bleep",
"blend",
"bless",
"blimp",
"blind",
"blink",
"blips",
"bliss",
"blitz",
"block",
"bloke",
"blond",
"blood",
"bloom",
"blown",
"blows",
"bluer",
"blues",
"bluff",
"blume",
"blunt",
"blurb",
"blurt",
"blush",
"board",
"boast",
"boats",
"bobby",
"bogus",
"boils",
"bolts",
"bombs",
"bonds",
"boned",
"boner",
"bones",
"bongo",
"bonus",
"boobs",
"booby",
"books",
"boost",
"booth",
"boots",
"booty",
"booze",
"bored",
"bosom",
"bossy",
"bound",
"bouts",
"bowed",
"bowel",
"bowls",
"boxed",
"boxer",
"boxes",
"bozos",
"brace",
"brags",
"braid",
"brain",
"brake",
"brand",
"brash",
"brass",
"brats",
"brava",
"brave",
"bravo",
"brawl",
"brays",
"bread",
"break",
"breed",
"brent",
"brews",
"briar",
"bribe",
"brick",
"bride",
"brief",
"brill",
"bring",
"brink",
"britt",
"broad",
"brock",
"broke",
"brood",
"brook",
"broom",
"broth",
"brown",
"brunt",
"brush",
"brute",
"bucko",
"bucks",
"buddy",
"budge",
"buffs",
"buffy",
"buggy",
"bugle",
"build",
"built",
"bulbs",
"bulge",
"bulky",
"bulls",
"bully",
"bumps",
"bumpy",
"bunch",
"bunks",
"bunny",
"burbs",
"burke",
"burly",
"burns",
"burnt",
"burro",
"burst",
"buses",
"busts",
"busty",
"butch",
"butts",
"buyer",
"bwana",
"cabin",
"cable",
"cache",
"caddy",
"cadet",
"caged",
"cages",
"cagey",
"cakes",
"calls",
"calms",
"camel",
"camps",
"canal",
"candy",
"canoe",
"caper",
"carat",
"carbo",
"carbs",
"cards",
"cared",
"cares",
"cargo",
"carol",
"carry",
"carts",
"carve",
"cased",
"cases",
"caste",
"casts",
"catch",
"cater",
"cates",
"catty",
"cause",
"caved",
"caves",
"cease",
"cedar",
"cello",
"cells",
"cents",
"chaff",
"chain",
"chair",
"chalk",
"champ",
"chang",
"chant",
"chaos",
"chaps",
"charm",
"chart",
"chase",
"chasm",
"cheap",
"cheat",
"check",
"cheek",
"cheep",
"cheer",
"chefs",
"chemo",
"chess",
"chest",
"chevy",
"chewy",
"chick",
"chico",
"chief",
"child",
"chile",
"chili",
"chill",
"chime",
"chimp",
"china",
"chink",
"chino",
"chins",
"chips",
"chirp",
"choir",
"choke",
"chomp",
"chops",
"chord",
"chore",
"chose",
"chuck",
"chump",
"chums",
"chunk",
"churn",
"chute",
"cider",
"cigar",
"cissy",
"cited",
"civic",
"civil",
"clack",
"claim",
"clamp",
"clams",
"clang",
"clash",
"clasp",
"class",
"claws",
"clean",
"clear",
"clerk",
"click",
"cliff",
"climb",
"cling",
"clink",
"clips",
"cloak",
"clock",
"clods",
"clogs",
"clone",
"close",
"cloth",
"clots",
"cloud",
"clout",
"clown",
"clubs",
"cluck",
"clues",
"clump",
"clung",
"clunk",
"coach",
"coals",
"coast",
"coats",
"cobra",
"cocky",
"cocoa",
"coded",
"codes",
"coeds",
"coins",
"cokes",
"colds",
"coles",
"colic",
"colin",
"colon",
"color",
"comas",
"combo",
"comer",
"comes",
"comet",
"comfy",
"comic",
"comma",
"condo",
"cones",
"coney",
"cooks",
"cools",
"coral",
"cords",
"corks",
"corky",
"corny",
"corps",
"costa",
"costs",
"couch",
"cough",
"could",
"count",
"coupe",
"court",
"coven",
"cover",
"covet",
"cowed",
"cower",
"crabs",
"crack",
"craft",
"cramp",
"crane",
"crank",
"craps",
"crash",
"crass",
"crate",
"crave",
"crawl",
"craze",
"crazy",
"creak",
"cream",
"credo",
"creed",
"creek",
"creep",
"creme",
"crepe",
"crept",
"crest",
"crews",
"cribs",
"crick",
"cried",
"crier",
"cries",
"crime",
"crimp",
"crisp",
"croak",
"crock",
"croft",
"crook",
"croon",
"crops",
"cross",
"crowd",
"crown",
"crows",
"crude",
"cruel",
"crumb",
"crush",
"crust",
"crypt",
"cubby",
"cubed",
"cubes",
"cubic",
"cuddy",
"cuffs",
"culpa",
"cumin",
"cunts",
"cupid",
"cuppa",
"cured",
"cures",
"curie",
"curly",
"curry",
"curse",
"curve",
"cushy",
"cuter",
"cutie",
"cycle",
"cynic",
"daddy",
"daffy",
"daily",
"dairy",
"daisy",
"dally",
"dance",
"dandy",
"dared",
"dares",
"darks",
"darts",
"dated",
"dater",
"dates",
"deals",
"dealt",
"dears",
"death",
"debit",
"debts",
"debut",
"decaf",
"decay",
"decks",
"decor",
"decoy",
"deeds",
"deets",
"deity",
"delay",
"delly",
"delta",
"delve",
"demon",
"demur",
"dense",
"dents",
"depot",
"depth",
"derby",
"desks",
"deuce",
"devil",
"devon",
"dials",
"diary",
"diced",
"dicey",
"dicks",
"dicky",
"diets",
"digit",
"dildo",
"dills",
"dilly",
"dimes",
"dimly",
"dined",
"diner",
"dingo",
"dings",
"dingy",
"dinks",
"dinky",
"dirty",
"disco",
"discs",
"disks",
"ditch",
"ditsy",
"ditto",
"ditty",
"divas",
"diver",
"dives",
"divvy",
"dizzy",
"docks",
"dodge",
"dodgy",
"doggy",
"doily",
"doing",
"dolce",
"dolls",
"dolly",
"domes",
"dongs",
"donna",
"donor",
"donut",
"doors",
"doozy",
"dopes",
"dopey",
"dorky",
"dorms",
"dosed",
"doses",
"dotes",
"dotty",
"doubt",
"dough",
"doves",
"dowdy",
"downs",
"downy",
"dozed",
"dozen",
"dozer",
"draft",
"drags",
"drain",
"drake",
"drama",
"drank",
"drape",
"drawn",
"draws",
"dread",
"dream",
"dreck",
"dregs",
"dress",
"dried",
"drier",
"dries",
"drift",
"drill",
"drink",
"drips",
"drive",
"droll",
"drone",
"drool",
"droop",
"drops",
"drove",
"drown",
"drugs",
"drums",
"drunk",
"dryer",
"ducks",
"ducky",
"ducts",
"dudes",
"dukes",
"dummy",
"dumps",
"dunes",
"dunks",
"duped",
"dusky",
"dusty",
"dutch",
"duvet",
"dwarf",
"dweeb",
"dwell",
"dying",
"dykes",
"eager",
"eagle",
"early",
"earns",
"earth",
"easel",
"eases",
"eaten",
"eater",
"eaves",
"ebony",
"edema",
"edged",
"edges",
"edict",
"edits",
"eerie",
"eight",
"eject",
"eking",
"elbow",
"elder",
"elect",
"elite",
"elope",
"elude",
"elves",
"ember",
"emery",
"empty",
"enact",
"ended",
"enema",
"enemy",
"enjoy",
"ennui",
"ensue",
"enter",
"entry",
"envoy",
"epoxy",
"equal",
"erase",
"erect",
"erica",
"erode",
"erred",
"error",
"erupt",
"essay",
"ester",
"ether",
"ethic",
"ethyl",
"euros",
"evade",
"evens",
"event",
"every",
"evict",
"evils",
"exact",
"exams",
"excel",
"execs",
"exile",
"exist",
"exits",
"expel",
"extra",
"fable",
"faced",
"faces",
"facet",
"facts",
"faded",
"fades",
"faggy",
"fails",
"faint",
"fairy",
"faith",
"faked",
"falls",
"false",
"famed",
"fancy",
"fangs",
"fanny",
"farce",
"fared",
"farms",
"farts",
"fatal",
"fates",
"fatso",
"fatty",
"fault",
"fauna",
"favor",
"faxed",
"faxes",
"fears",
"feast",
"feats",
"feces",
"feeds",
"feels",
"feign",
"fella",
"felon",
"femme",
"femur",
"fence",
"ferry",
"fetal",
"fetch",
"fetus",
"feuds",
"fever",
"fewer",
"fiber",
"fibre",
"ficus",
"field",
"fiend",
"fiery",
"fifth",
"fifty",
"fight",
"filed",
"files",
"filet",
"fills",
"filly",
"films",
"filth",
"final",
"finch",
"finds",
"fined",
"finer",
"fines",
"fired",
"fires",
"firms",
"first",
"fishy",
"fists",
"fitch",
"fiver",
"fives",
"fixed",
"fixer",
"fixes",
"flack",
"flags",
"flail",
"flair",
"flake",
"flaky",
"flame",
"flank",
"flaps",
"flare",
"flash",
"flask",
"flats",
"flaws",
"fleas",
"fleet",
"flesh",
"flick",
"flier",
"flies",
"fling",
"flint",
"flips",
"flirt",
"float",
"flock",
"flood",
"floor",
"flops",
"flora",
"floss",
"flour",
"flown",
"flows",
"fluff",
"fluid",
"fluke",
"flung",
"flunk",
"flush",
"flute",
"flyer",
"foamy",
"focal",
"focus",
"folds",
"folks",
"folly",
"foods",
"fools",
"foray",
"force",
"forge",
"forgo",
"forks",
"forms",
"forte",
"forth",
"forty",
"forum",
"fosse",
"found",
"fours",
"foxes",
"foyer",
"frail",
"frame",
"franc",
"frank",
"fraud",
"freak",
"freed",
"freer",
"frees",
"fresh",
"fried",
"fries",
"frisk",
"fritz",
"frogs",
"front",
"frost",
"frown",
"froze",
"fruit",
"fryer",
"fucks",
"fudge",
"fuels",
"fugue",
"fully",
"fumes",
"funds",
"fungi",
"funky",
"funny",
"furry",
"fused",
"fuses",
"fussy",
"futon",
"fuzzy",
"gabby",
"gains",
"gamer",
"games",
"gamma",
"gammy",
"gamut",
"gangs",
"ganja",
"garth",
"gases",
"gassy",
"gated",
"gates",
"gator",
"gaudy",
"gauge",
"gault",
"gauze",
"gavel",
"gazed",
"gears",
"gecko",
"geeks",
"geeky",
"geese",
"gemma",
"genes",
"genie",
"genoa",
"genre",
"gents",
"genus",
"germs",
"getup",
"ghost",
"ghoul",
"giant",
"giddy",
"gifts",
"gills",
"gimme",
"ginny",
"girls",
"girly",
"girth",
"given",
"giver",
"gives",
"gizmo",
"glade",
"gland",
"glare",
"glass",
"glaze",
"gleam",
"glide",
"glint",
"gloat",
"globe",
"gloom",
"glory",
"gloss",
"glove",
"glued",
"glues",
"gnats",
"gnome",
"goals",
"goats",
"godly",
"goers",
"gofer",
"going",
"golly",
"goner",
"gonzo",
"goods",
"goody",
"gooey",
"goofy",
"goons",
"goopy",
"goose",
"gorge",
"gouge",
"gourd",
"gowns",
"grabs",
"grace",
"grade",
"grail",
"grain",
"grams",
"grand",
"grant",
"grape",
"grasp",
"grass",
"grate",
"grave",
"gravy",
"great",
"greed",
"greek",
"green",
"greet",
"grief",
"griff",
"grift",
"grill",
"grime",
"grind",
"grins",
"gripe",
"grips",
"grits",
"groin",
"groom",
"grope",
"gross",
"group",
"grove",
"growl",
"grown",
"grows",
"grubs",
"gruel",
"grunt",
"guard",
"guava",
"guess",
"guest",
"guide",
"guild",
"guilt",
"gulag",
"gulch",
"gushy",
"gusto",
"gypsy",
"habit",
"hacks",
"hades",
"hails",
"hairs",
"hairy",
"hallo",
"halls",
"hands",
"handy",
"hangs",
"hanks",
"hanky",
"happy",
"hardy",
"harem",
"harms",
"harps",
"harpy",
"harry",
"harsh",
"harts",
"haste",
"hasty",
"hatch",
"hated",
"hater",
"hates",
"hauls",
"haunt",
"haute",
"haven",
"havoc",
"hawks",
"hazel",
"heads",
"heady",
"heals",
"heaps",
"heard",
"hears",
"heart",
"heath",
"heats",
"heave",
"heavy",
"hedge",
"heels",
"hefty",
"heigh",
"heirs",
"heist",
"helix",
"hello",
"hells",
"helms",
"helps",
"hence",
"henry",
"herbs",
"herds",
"hertz",
"hexes",
"hicks",
"hides",
"highs",
"hiked",
"hiker",
"hills",
"hints",
"hippo",
"hired",
"hires",
"hitch",
"hives",
"hobby",
"hocks",
"hocus",
"hogan",
"hoist",
"hokey",
"holds",
"holed",
"holes",
"holly",
"homer",
"homes",
"homey",
"homos",
"honda",
"honed",
"honey",
"honks",
"honky",
"honor",
"hooch",
"hooks",
"hooky",
"hoops",
"hoped",
"hopes",
"hoppy",
"horde",
"horns",
"horny",
"horse",
"hosed",
"hoses",
"hosts",
"hotel",
"hound",
"hours",
"house",
"hovel",
"hover",
"howdy",
"hubby",
"huffy",
"hullo",
"human",
"humid",
"humor",
"humph",
"humps",
"humus",
"hunch",
"hunks",
"hunky",
"hurry",
"hurst",
"hurts",
"husks",
"husky",
"hussy",
"hutch",
"hydra",
"hymns",
"hyped",
"hyper",
"icing",
"icons",
"ideal",
"ideas",
"idiom",
"idiot",
"idols",
"igloo",
"iliad",
"image",
"imply",
"incur",
"index",
"indie",
"inept",
"infra",
"inner",
"input",
"inter",
"intro",
"irons",
"irony",
"issue",
"itchy",
"items",
"ivory",
"jabot",
"jacks",
"jaded",
"japan",
"jaunt",
"jawed",
"jeans",
"jeeps",
"jelly",
"jenny",
"jerks",
"jerky",
"jerry",
"jesse",
"jewel",
"jiffy",
"jihad",
"jimmy",
"jocks",
"johns",
"joins",
"joint",
"joked",
"joker",
"jokes",
"jolly",
"jones",
"joust",
"judas",
"judge",
"juice",
"juicy",
"julep",
"jumbo",
"jumps",
"jumpy",
"junky",
"juror",
"kabob",
"kappa",
"kaput",
"karat",
"karma",
"kasha",
"kayak",
"keeps",
"kelly",
"kendo",
"kerry",
"ketch",
"khaki",
"kicks",
"kicky",
"kiddo",
"kills",
"kilos",
"kinds",
"kings",
"kinky",
"kiosk",
"kissy",
"kites",
"kitty",
"klutz",
"knack",
"kneed",
"kneel",
"knees",
"knelt",
"knife",
"knobs",
"knock",
"knoll",
"knots",
"known",
"knows",
"koala",
"kooks",
"kooky",
"kraft",
"kudos",
"label",
"labor",
"laced",
"laces",
"lacey",
"lacks",
"laden",
"ladle",
"lager",
"laird",
"laker",
"lakes",
"lambs",
"lamps",
"lance",
"lands",
"lanes",
"lanky",
"lapel",
"lapse",
"larch",
"large",
"laser",
"lasso",
"lasts",
"latch",
"later",
"latex",
"lathe",
"latte",
"laugh",
"laura",
"lawns",
"layer",
"lazar",
"leads",
"leafs",
"leafy",
"leaks",
"leaky",
"leans",
"leaps",
"leapt",
"learn",
"leary",
"lease",
"leash",
"least",
"leave",
"ledge",
"leech",
"leery",
"lefts",
"lefty",
"legal",
"leggy",
"legit",
"lemon",
"lemur",
"leper",
"levee",
"level",
"lever",
"lewis",
"liane",
"liars",
"libel",
"licks",
"liege",
"lifer",
"lifts",
"light",
"liked",
"likes",
"lilac",
"limbo",
"limbs",
"limes",
"limey",
"limit",
"limos",
"limps",
"lined",
"linen",
"liner",
"lines",
"lingo",
"links",
"lions",
"lippy",
"lists",
"liter",
"lived",
"liven",
"liver",
"lives",
"livid",
"llama",
"loads",
"loans",
"loath",
"lobby",
"lobes",
"local",
"locks",
"lodge",
"lofty",
"logan",
"logic",
"loins",
"lolly",
"loner",
"longs",
"looks",
"loons",
"loony",
"loops",
"loose",
"loran",
"lords",
"lorry",
"loser",
"loses",
"lotte",
"lotto",
"louie",
"louis",
"louse",
"lousy",
"loved",
"lover",
"loves",
"lower",
"lowly",
"loyal",
"lucid",
"lucky",
"lumps",
"lumpy",
"lunar",
"lunch",
"lunge",
"lungs",
"lupus",
"lurch",
"lured",
"lures",
"lurks",
"lusts",
"lying",
"lymph",
"lynch",
"mache",
"macho",
"madam",
"madly",
"madre",
"mafia",
"magic",
"maids",
"mails",
"major",
"maker",
"makes",
"males",
"malls",
"mamie",
"mamma",
"mangy",
"manic",
"manly",
"manna",
"manor",
"maple",
"march",
"marge",
"maria",
"marks",
"marry",
"marsh",
"masks",
"mason",
"massa",
"match",
"mated",
"mates",
"matey",
"mavis",
"maxim",
"mayan",
"maybe",
"mayor",
"meals",
"mealy",
"means",
"meant",
"meats",
"meaty",
"mecca",
"medal",
"media",
"medic",
"meets",
"melon",
"melts",
"memos",
"mensa",
"menus",
"mercy",
"merit",
"merle",
"merry",
"messy",
"metal",
"meter",
"metro",
"micro",
"midge",
"midst",
"miggs",
"might",
"mikes",
"miles",
"milky",
"mills",
"mimes",
"mimic",
"mince",
"minds",
"mined",
"miner",
"mines",
"minks",
"minor",
"mints",
"minty",
"minus",
"mirth",
"missy",
"misty",
"mites",
"mitts",
"mixed",
"mixer",
"mixes",
"mixup",
"moans",
"mocha",
"mocks",
"model",
"modem",
"modus",
"mohel",
"moist",
"molds",
"moldy",
"moles",
"molly",
"molto",
"momma",
"mommy",
"monde",
"mondo",
"money",
"monks",
"monte",
"month",
"mooch",
"moods",
"moody",
"moola",
"moons",
"moors",
"moose",
"moped",
"mopes",
"mopey",
"moral",
"moron",
"morph",
"morse",
"mosey",
"motel",
"moths",
"motif",
"motor",
"motto",
"mould",
"mound",
"mount",
"mourn",
"mouse",
"mousy",
"mouth",
"moved",
"moves",
"movie",
"mowed",
"mower",
"moxie",
"mucus",
"muddy",
"muggy",
"mulch",
"mules",
"muley",
"mummy",
"mumps",
"munch",
"mural",
"murky",
"muses",
"mushy",
"music",
"myths",
"nacho",
"nails",
"naive",
"naked",
"named",
"names",
"nance",
"nancy",
"nanny",
"nappy",
"nasal",
"nasty",
"natty",
"naval",
"nazis",
"necks",
"needs",
"needy",
"nelly",
"nerds",
"nerdy",
"nerve",
"never",
"newly",
"nexus",
"nicer",
"niche",
"nicks",
"niece",
"nifty",
"night",
"nines",
"ninja",
"ninny",
"ninth",
"nippy",
"nitty",
"nixed",
"noble",
"nodes",
"noise",
"noisy",
"nomad",
"noose",
"north",
"nosed",
"noses",
"nosey",
"notch",
"noted",
"notes",
"novel",
"nudes",
"nudge",
"nudie",
"nuked",
"nukes",
"nurse",
"nutty",
"nylon",
"nymph",
"oasis",
"oaths",
"obese",
"obits",
"occur",
"ocean",
"oddly",
"offer",
"often",
"ogres",
"oiled",
"olden",
"older",
"oldie",
"olive",
"omega",
"omens",
"onion",
"oomph",
"opens",
"opera",
"opium",
"opted",
"optic",
"orbed",
"orbit",
"order",
"organ",
"other",
"otter",
"ought",
"ounce",
"outdo",
"outer",
"owing",
"owned",
"owner",
"ozone",
"paced",
"pacer",
"paces",
"packs",
"paddy",
"padre",
"pagan",
"paged",
"pager",
"pages",
"pains",
"paint",
"pairs",
"paler",
"pales",
"palms",
"palsy",
"panda",
"panel",
"panic",
"pansy",
"pants",
"panty",
"paper",
"pappy",
"paris",
"parka",
"parks",
"parts",
"party",
"pasta",
"paste",
"pasts",
"patch",
"paths",
"patio",
"patsy",
"patty",
"pause",
"paved",
"peace",
"peach",
"peaks",
"pearl",
"pecan",
"pecks",
"pedal",
"pedro",
"peeks",
"peels",
"peeps",
"peers",
"pelts",
"penal",
"pence",
"penis",
"penne",
"penny",
"perch",
"peril",
"perks",
"perky",
"perry",
"pesky",
"pesos",
"pesto",
"pests",
"petal",
"peter",
"petit",
"petty",
"phase",
"phone",
"phony",
"photo",
"piano",
"picks",
"picky",
"piece",
"piggy",
"pilar",
"piled",
"piles",
"pills",
"pilot",
"pinch",
"pines",
"pinks",
"pinky",
"pinot",
"pinto",
"pints",
"pious",
"piper",
"pipes",
"pitch",
"pithy",
"pivot",
"pixie",
"pizza",
"place",
"plaid",
"plain",
"plait",
"plane",
"plank",
"plans",
"plant",
"plate",
"playa",
"plays",
"plaza",
"plead",
"pleas",
"plots",
"pluck",
"plugs",
"plump",
"plums",
"plush",
"poach",
"poems",
"poets",
"point",
"poise",
"poked",
"poker",
"pokes",
"pokey",
"polar",
"poles",
"polio",
"polka",
"polls",
"ponds",
"pooch",
"poofs",
"poofy",
"pools",
"poppa",
"poppy",
"porch",
"pores",
"porky",
"porno",
"ports",
"posed",
"poser",
"poses",
"posse",
"posts",
"potty",
"pouch",
"pound",
"pours",
"power",
"prank",
"prays",
"press",
"preys",
"price",
"prick",
"pride",
"prima",
"prime",
"primo",
"print",
"prior",
"priss",
"privy",
"prize",
"probe",
"promo",
"proms",
"prone",
"proof",
"props",
"prose",
"proud",
"prove",
"prowl",
"proxy",
"prude",
"prune",
"psalm",
"psych",
"pubes",
"pubic",
"puffs",
"puffy",
"pulls",
"pulse",
"pumps",
"punch",
"punks",
"punky",
"pupil",
"puppy",
"puree",
"purer",
"purge",
"purse",
"pushy",
"pussy",
"putty",
"pygmy",
"quack",
"quake",
"quark",
"quart",
"queen",
"queer",
"quell",
"query",
"quest",
"queue",
"quick",
"quiet",
"quilt",
"quirk",
"quite",
"quits",
"quota",
"quote",
"quoth",
"rabbi",
"rabid",
"raced",
"racer",
"races",
"racks",
"radar",
"radio",
"rages",
"raids",
"rails",
"rains",
"rainy",
"raise",
"rajah",
"raked",
"rally",
"ralph",
"ramus",
"ranch",
"randy",
"range",
"ranks",
"rants",
"raped",
"rapes",
"rapid",
"rated",
"rates",
"ratio",
"ratty",
"raved",
"raven",
"rayed",
"razor",
"reach",
"react",
"reads",
"ready",
"realm",
"rears",
"rebel",
"recap",
"recon",
"redid",
"reefs",
"reeks",
"reels",
"reeve",
"refer",
"regal",
"rehab",
"reign",
"relax",
"relay",
"relic",
"renal",
"renew",
"rents",
"repay",
"reply",
"reset",
"resin",
"rests",
"retro",
"rhino",
"rhyme",
"ricks",
"rider",
"rides",
"ridge",
"rifle",
"right",
"rigid",
"rigor",
"riled",
"riley",
"rings",
"rinse",
"rioja",
"riots",
"risen",
"rises",
"risks",
"risky",
"rites",
"ritzy",
"rival",
"river",
"roach",
"roads",
"roast",
"robes",
"robin",
"robot",
"rocks",
"rocky",
"rodeo",
"roger",
"rogue",
"roles",
"rolls",
"roman",
"romeo",
"roofs",
"rooms",
"roomy",
"roost",
"roots",
"roped",
"ropes",
"roses",
"rosin",
"rouge",
"rough",
"round",
"rouse",
"roust",
"route",
"rover",
"rowan",
"rowdy",
"royal",
"rubes",
"ruder",
"rugby",
"ruins",
"ruled",
"ruler",
"rules",
"rumba",
"rummy",
"rumor",
"runes",
"runny",
"rural",
"rusty",
"saber",
"sabin",
"sable",
"sacks",
"sadly",
"safer",
"safes",
"sahib",
"sails",
"saint",
"saith",
"sakes",
"salad",
"sales",
"sally",
"salon",
"salsa",
"salty",
"sands",
"sandy",
"santo",
"sappy",
"saran",
"sarge",
"sassy",
"satin",
"satyr",
"sauce",
"sauna",
"saved",
"saver",
"saves",
"savin",
"savor",
"savvy",
"sawed",
"sayer",
"scabs",
"scald",
"scale",
"scalp",
"scamp",
"scams",
"scans",
"scant",
"scare",
"scarf",
"scars",
"scary",
"scene",
"scent",
"schmo",
"scoff",
"scone",
"scoop",
"scoot",
"scope",
"score",
"scots",
"scout",
"scram",
"scrap",
"screw",
"scrub",
"scuba",
"scuff",
"seals",
"seams",
"sears",
"seats",
"sedan",
"seeds",
"seedy",
"seeks",
"seems",
"segue",
"seize",
"sells",
"semen",
"sends",
"senor",
"sense",
"serge",
"serum",
"serve",
"setup",
"seven",
"sever",
"sewed",
"sewer",
"sexes",
"shack",
"shade",
"shady",
"shaft",
"shake",
"shaky",
"shale",
"shall",
"shalt",
"shame",
"shank",
"shape",
"share",
"shark",
"sharp",
"shave",
"shawl",
"shawn",
"shear",
"sheds",
"sheen",
"sheep",
"sheer",
"sheet",
"shelf",
"shell",
"shift",
"shill",
"shine",
"shins",
"shiny",
"ships",
"shirt",
"shits",
"shiva",
"shoal",
"shock",
"shoes",
"shone",
"shook",
"shoot",
"shops",
"shore",
"short",
"shots",
"shout",
"shove",
"shown",
"shows",
"showy",
"shred",
"shrew",
"shrub",
"shrug",
"shuck",
"shunt",
"shush",
"shuts",
"sicko",
"sided",
"sides",
"sidle",
"siege",
"sighs",
"sight",
"sigma",
"signs",
"sikes",
"silks",
"silky",
"silly",
"silva",
"since",
"singe",
"sings",
"sinks",
"sinus",
"siree",
"siren",
"sissy",
"sites",
"sixes",
"sixth",
"sixty",
"sized",
"sizes",
"skate",
"skids",
"skied",
"skier",
"skies",
"skiff",
"skill",
"skimp",
"skins",
"skirt",
"skulk",
"skull",
"skunk",
"slack",
"slams",
"slang",
"slant",
"slaps",
"slash",
"slate",
"slave",
"sleek",
"sleep",
"sleet",
"slept",
"slice",
"slick",
"slide",
"slime",
"slimy",
"sling",
"slink",
"slips",
"slope",
"slots",
"slugs",
"slung",
"slurp",
"slush",
"sluts",
"smack",
"small",
"smart",
"smash",
"smear",
"smell",
"smelt",
"smile",
"smirk",
"smite",
"smith",
"smock",
"smoke",
"smoky",
"snack",
"snags",
"snake",
"snaps",
"snarl",
"sneak",
"sneer",
"snide",
"sniff",
"snipe",
"snook",
"snoop",
"snore",
"snort",
"snout",
"snowy",
"snuck",
"snuff",
"soaps",
"soapy",
"soars",
"sober",
"socks",
"sodas",
"sofas",
"softy",
"soggy",
"solar",
"soles",
"solid",
"solve",
"sonar",
"songs",
"sonny",
"sorel",
"sores",
"sorry",
"sorts",
"souls",
"sound",
"soups",
"soupy",
"souse",
"south",
"space",
"spade",
"spank",
"spans",
"spare",
"spark",
"spasm",
"spate",
"spawn",
"speak",
"spear",
"speck",
"specs",
"speed",
"spell",
"spelt",
"spend",
"spent",
"sperm",
"spice",
"spicy",
"spied",
"spiel",
"spies",
"spike",
"spiky",
"spill",
"spine",
"spins",
"spiny",
"spite",
"spits",
"spitz",
"splat",
"split",
"spoil",
"spoke",
"spook",
"spool",
"spoon",
"sport",
"spots",
"spray",
"spree",
"spunk",
"spurs",
"spurt",
"squad",
"squat",
"squaw",
"squid",
"stack",
"staff",
"stage",
"stain",
"stair",
"stake",
"stale",
"stalk",
"stall",
"stamp",
"stand",
"stang",
"stare",
"stark",
"stars",
"start",
"stash",
"state",
"stats",
"stave",
"stays",
"stead",
"steak",
"steal",
"steam",
"steed",
"steel",
"steep",
"steer",
"stein",
"stems",
"steno",
"steps",
"stern",
"stick",
"stiff",
"still",
"sting",
"stink",
"stint",
"stirs",
"stock",
"stoic",
"stoke",
"stole",
"stomp",
"stone",
"stony",
"stood",
"stool",
"stoop",
"stops",
"store",
"stork",
"storm",
"story",
"stove",
"strap",
"straw",
"stray",
"strep",
"strip",
"strut",
"stubs",
"stuck",
"studs",
"study",
"stuff",
"stump",
"stung",
"stunk",
"stunt",
"style",
"suave",
"sucks",
"suede",
"sugar",
"suing",
"suite",
"suits",
"sunny",
"super",
"surge",
"surly",
"sushi",
"sutra",
"swabs",
"swami",
"swamp",
"swank",
"swans",
"swarm",
"swear",
"sweat",
"swede",
"sweep",
"sweet",
"swell",
"swept",
"swift",
"swill",
"swims",
"swine",
"swing",
"swipe",
"swirl",
"swiss",
"swoop",
"sword",
"swore",
"sworn",
"swung",
"sykes",
"synch",
"syrup",
"tabby",
"table",
"tacit",
"tacks",
"tacky",
"tacos",
"taffy",
"tails",
"taint",
"taken",
"taker",
"takes",
"takin",
"tales",
"talks",
"talky",
"tally",
"tammy",
"tango",
"tanks",
"tapas",
"taped",
"tapes",
"tardy",
"tarot",
"tarts",
"tasks",
"taste",
"tasty",
"tater",
"taunt",
"taxed",
"taxes",
"taxis",
"teach",
"teams",
"tears",
"teary",
"tease",
"teddy",
"teens",
"teeny",
"teeth",
"telex",
"tells",
"telly",
"tempo",
"temps",
"tempt",
"tends",
"tenor",
"tense",
"tenth",
"tents",
"tepid",
"terms",
"terra",
"terry",
"tests",
"testy",
"texas",
"texts",
"thank",
"theft",
"their",
"theme",
"there",
"these",
"theta",
"thick",
"thief",
"thigh",
"thine",
"thing",
"think",
"thins",
"third",
"thong",
"thorn",
"those",
"three",
"threw",
"throw",
"thugs",
"thumb",
"thump",
"thyme",
"tiara",
"tibia",
"ticks",
"tidal",
"tides",
"tiger",
"tight",
"tiles",
"timed",
"timer",
"times",
"timid",
"tippy",
"tipsy",
"tired",
"tires",
"titan",
"title",
"titty",
"tizzy",
"toast",
"today",
"toddy",
"token",
"tolls",
"tombs",
"tommy",
"toned",
"toner",
"tones",
"tongs",
"tonic",
"tools",
"toons",
"tooth",
"toots",
"topaz",
"topes",
"topic",
"torah",
"torch",
"torso",
"total",
"totem",
"touch",
"tough",
"tours",
"towed",
"towel",
"tower",
"towns",
"toxic",
"toxin",
"toyed",
"trace",
"track",
"trade",
"trail",
"train",
"trait",
"tramp",
"trans",
"traps",
"trash",
"trays",
"tread",
"treat",
"trees",
"trend",
"triad",
"trial",
"tribe",
"trick",
"tried",
"tries",
"tripe",
"trips",
"trite",
"troll",
"troop",
"trout",
"trove",
"truce",
"truck",
"truer",
"truly",
"trump",
"trunk",
"truss",
"trust",
"truth",
"tubby",
"tubes",
"tulip",
"tulle",
"tummy",
"tumor",
"tuned",
"tunes",
"tunic",
"turbo",
"turds",
"turks",
"turns",
"tushy",
"tutor",
"tutti",
"tuxes",
"twain",
"tweak",
"tweed",
"tween",
"tweet",
"twerp",
"twice",
"twigs",
"twine",
"twins",
"twirl",
"twist",
"twits",
"tying",
"tykes",
"typed",
"types",
"tyres",
"ulcer",
"ultra",
"uncle",
"uncut",
"under",
"undue",
"unfit",
"union",
"unite",
"units",
"unity",
"untie",
"until",
"unzip",
"upped",
"upper",
"upset",
"urban",
"urged",
"urges",
"urine",
"users",
"usher",
"using",
"usual",
"utter",
"vague",
"valet",
"valid",
"valor",
"value",
"valve",
"vamps",
"vapid",
"vault",
"vegan",
"veils",
"veins",
"venom",
"vents",
"venue",
"verbs",
"verge",
"verse",
"vesta",
"vests",
"vials",
"vibes",
"vicar",
"video",
"views",
"vigil",
"vigor",
"villa",
"vinyl",
"viola",
"viper",
"viral",
"virus",
"visas",
"visit",
"visor",
"vista",
"vital",
"vivid",
"vixen",
"vocal",
"vodka",
"vogue",
"voice",
"voila",
"volts",
"vomit",
"voted",
"voter",
"votes",
"vouch",
"vowed",
"vowel",
"vroom",
"vying",
"wacko",
"wacky",
"wager",
"wages",
"wagon",
"wahoo",
"waist",
"waits",
"waive",
"waken",
"wakes",
"walks",
"walla",
"walls",
"wally",
"waltz",
"wants",
"wares",
"warms",
"warts",
"washy",
"wasps",
"waste",
"watch",
"water",
"watts",
"waved",
"waves",
"waxed",
"waxes",
"wears",
"weary",
"weave",
"weber",
"wedge",
"weeds",
"weeks",
"weeny",
"weepy",
"weigh",
"weird",
"welch",
"wells",
"welsh",
"welts",
"wench",
"whack",
"whale",
"wharf",
"whats",
"wheat",
"wheel",
"where",
"which",
"whiff",
"while",
"whims",
"whine",
"whiny",
"whirl",
"whisk",
"white",
"whole",
"whoop",
"whore",
"whose",
"widen",
"wider",
"widow",
"width",
"wield",
"wiggy",
"wills",
"willy",
"wimps",
"wimpy",
"winch",
"winds",
"windy",
"wings",
"winks",
"winos",
"wiped",
"wiper",
"wipes",
"wired",
"wires",
"wised",
"wiser",
"witch",
"witty",
"wives",
"woken",
"woman",
"women",
"wonky",
"woods",
"woody",
"wooed",
"woops",
"woozy",
"words",
"works",
"world",
"worms",
"worry",
"worse",
"worst",
"worth",
"would",
"wound",
"woven",
"wowed",
"wraps",
"wrath",
"wreak",
"wreck",
"wring",
"wrist",
"write",
"wrong",
"wrote",
"wrung",
"wussy",
"xerox",
"yacht",
"yahoo",
"yanks",
"yards",
"years",
"yeast",
"yells",
"yield",
"yikes",
"yodel",
"yokel",
"young",
"yours",
"youse",
"youth",
"yummy",
"zebra",
"zeros",
"zesty",
"zippy",
"zoned",
"zones",
}

func verifyGuess(guess, word string) (bool, string){
	var builder strings.Builder
	builder.WriteString("-")
	for i := 0; i < len(word); i++ {
		c1 := word[i]
		c2 := guess[i]

		if c1 == c2 {
			builder.WriteString(fmt.Sprintf("%v%s%v-", Green, string(c2), Reset))
		} else if strings.Contains(word, string(c2)) {
			builder.WriteString(fmt.Sprintf("%v%s%v-", Magenta, string(c2), Reset))
		} else {
			builder.WriteString(fmt.Sprintf("%v%s%v-", Red, string(c2), Reset))
		}
	}
	if guess == word {
		return true, builder.String()
	}

	return false, builder.String()
}

func buildTop() string {
	return fmt.Sprintf("%v───────────%v\n", tl, tr)
}

func buildBottom() string {
	return fmt.Sprintf("%v───────────%v", bl, br)
}

func gameBoard(userGuesses []string) string {
	var builder strings.Builder
	builder.WriteString(buildTop())
	
	for _, v := range userGuesses {
		builder.WriteString(fmt.Sprintf("│%v│\n", v))
	}


	builder.WriteString(buildBottom())
	return builder.String()
}

func gameStart() {
	clearTerminal()
	help()
	fmt.Println(gameBoard(nil))
}

func main() {
	word, err := getRandomWord()
	if err != nil {
		log.Fatal(err)
	}

	userGuesses := []string{}
	var guess string
	gameStart()


	for {
		fmt.Print("> ")
		fmt.Scan(&guess)
		if guess == ":q" {
			fmt.Printf("\nThe word was '%v%v%v'",Green, word, Reset)
			break
		}

		if len(guess) != 5{
			fmt.Printf("%vinvalid guess%v: want 5 have %v letters\n", Red, Reset, len(guess))
			continue
		}
		clearTerminal()

		ok, parsedstring := verifyGuess(guess, word)
		userGuesses = append(userGuesses, parsedstring)
		help()

		fmt.Println(gameBoard(userGuesses))

		if ok {
			fmt.Println(Green + "Congratulations!!!" + Reset)
			break
		}
	}
}