package main

import (
	"fmt"
	"io"
	"os"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	flag "github.com/spf13/pflag"
)

const (
	MIDDLE_EARTH = "(bil|bal|ban|hil|ham|hal|hol|hob|wil|me|or|ol|od|gor|for|fos|tol|ar|fin|ere|leo|vi|bi|bren|thor)" +
		"(|go|orbis|apol|adur|mos|ri|i|na|ole|n)" +
		"(|tur|axia|and|bo|gil|bin|bras|las|mac|grim|wise|l|lo|fo|co|ra|via|" +
		"da|ne|ta|y|wen|thiel|phin|dir|dor|tor|rod|on|rdo|dis)"
	JAPANESE_NAMES_CONSTRAINED = "(aka|aki|bashi|gawa|kawa|furu|fuku|fuji|hana|hara|haru|hashi|hira|hon|hoshi|" +
		"ichi|iwa|kami|kawa|ki|kita|kuchi|kuro|marui|matsu|miya|mori|moto|mura|nabe|naka|nishi|no|da|ta|o|oo|oka|" +
		"saka|saki|sawa|shita|shima|i|suzu|taka|take|to|toku|toyo|ue|wa|wara|wata|yama|yoshi|kei|ko|zawa|zen|sen|" +
		"ao|gin|kin|ken|shiro|zaki|yuki|asa)(||||||||||bashi|gawa|kawa|furu|fuku|fuji|hana|hara|haru|hashi|hira|" +
		"hon|hoshi|chi|wa|ka|kami|kawa|ki|kita|kuchi|kuro|marui|matsu|miya|mori|moto|mura|nabe|naka|nishi|no|da|" +
		"ta|o|oo|oka|saka|saki|sawa|shita|shima|suzu|taka|take|to|toku|toyo|ue|wa|wara|wata|yama|yoshi|kei|ko|" +
		"zawa|zen|sen|ao|gin|kin|ken|shiro|zaki|yuki|sa)"
	JAPANESE_NAMES_DIVERSE = "(a|i|u|e|o|||||)" +
		"(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|chi|tsu|te|to|na|ni|ni|nu|nu|ne|" +
		"no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|" +
		"re|ro|ro|ro|wa|wa|wa|wa|wo|wo)(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|" +
		"chi|tsu|te|to|na|ni|ni|nu|nu|ne|no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|" +
		"yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|re|ro|ro|ro|wa|wa|wa|wa|wo|wo)" +
		"(|(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|chi|tsu|te|to|na|ni|ni|nu|nu|ne|" +
		"no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|re|" +
		"ro|ro|ro|wa|wa|wa|wa|wo|wo)|(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|shi|su|su|se|so|ta|ta|chi|chi|" +
		"tsu|te|to|na|ni|ni|nu|nu|ne|no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|mu|mu|mu|me|mo|mo|mo|ya|yu|yu|" +
		"yu|yo|ra|ra|ra|ri|ru|ru|ru|re|ro|ro|ro|wa|wa|wa|wa|wo|wo)(|(ka|ki|ki|ku|ku|ke|ke|ko|ko|sa|sa|sa|shi|shi|" +
		"shi|su|su|se|so|ta|ta|chi|chi|tsu|te|to|na|ni|ni|nu|nu|ne|no|no|ha|hi|fu|fu|he|ho|ma|ma|ma|mi|mi|mi|mu|" +
		"mu|mu|mu|me|mo|mo|mo|ya|yu|yu|yu|yo|ra|ra|ra|ri|ru|ru|ru|re|ro|ro|ro|wa|wa|wa|wa|wo|wo)))(|||n)"
	CHINESE_NAMES = "(zh|x|q|sh|h)(ao|ian|uo|ou|ia)(|(l|w|c|p|b|m)(ao|ian|uo|ou|ia)" +
		"(|n)|-(l|w|c|p|b|m)(ao|ian|uo|ou|ia)(|(d|j|q|l)(a|ai|iu|ao|i)))"
	GREEK_NAMES      = "<s<v|V>(tia)|s<v|V>(os)|B<v|V>c(ios)|B<v|V><c|C>v(ios|os)>"
	HAWAIIAN_NAMES_1 = "((h|k|l|m|n|p|w)|)(a|e|i|o|u)((h|k|l|m|n|p|w)|)(a|e|i|o|u)(((h|k|l|m|n|p|w)|)" +
		"(a|e|i|o|u)|)(((h|k|l|m|n|p|w)|)(a|e|i|o|u)|)(((h|k|l|m|n|p|w)|)(a|e|i|o|u)|)(((h|k|l|m|n|p|w)|)" +
		"(a|e|i|o|u)|)"
	HAWAIIAN_NAMES_2 = "((h|k|l|m|n|p|w|)(a|e|i|o|u|a|e|i|o|u|ae|ai|ao|au|oi|ou|eu|ei)" +
		"(k|l|m|n|p|)|)(h|k|l|m|n|p|w|)(a|e|i|o|u|a|e|i|o|u|ae|ai|ao|au|oi|ou|eu|ei)(k|l|m|n|p|)"
	OLD_LATIN_PLACE_NAMES = "sv(nia|lia|cia|sia)"
	DRAGONS_PERN          = "<<s|ss>|<VC|vC|B|BVs|Vs>><v|V|v|<v(l|n|r)|vc>>(th)"
	DRAGON_RIDERS         = "c<s|cvc>"
	POKEMON               = "<i|s>v(mon|chu|zard|rtle)"
	FANTASY_VOWELS_R      = "(|(<B>|s|h|ty|ph|r))(i|ae|ya|ae|eu|ia|i|eo|ai|a)" +
		"(lo|la|sri|da|dai|the|sty|lae|due|li|lly|ri|na|ral|sur|rith)(|(su|nu|sti|llo|ria|))" +
		"(|(n|ra|p|m|lis|cal|deu|dil|suir|phos|ru|dru|rin|raap|rgue))"
	FANTASY_S_A = "(cham|chan|jisk|lis|frich|isk|lass|mind|sond|sund|ass|chad|lirt|und|mar|lis|il|<BVC>)" +
		"(jask|ast|ista|adar|irra|im|ossa|assa|osia|ilsa|<vCv>)(|(an|ya|la|sta|sda|sya|st|nya))"
	FANTASY_H_L = "(ch|cht|sh|cal|val|ell|har|shar|shal|rel|laen|ral|jht|alr|ch|cht|av)" +
		"(|(is|al|ow|ish|ul|el|ar|iel))" +
		"(aren|aeish|aith|even|adur|ulash|alith|atar|aia|erin|aera|ael|ira|iel|ahur|ishul)"
	FANTASY_N_L = "(ethr|qil|mal|er|eal|far|fil|fir|ing|ind|il|lam|quel|quar|quan|qar|pal|mal|yar|um|ard|enn|ey)" +
		"(|(<vc>|on|us|un|ar|as|en|ir|ur|at|ol|al|an))" +
		"(uard|wen|arn|on|il|ie|on|iel|rion|rian|an|ista|rion|rian|cil|mol|yon)"
	FANTASY_K_N = "(taith|kach|chak|kank|kjar|rak|kan|kaj|tach|rskal|kjol|jok|jor|jad|kot|kon|" +
		"knir|kror|kol|tul|rhaok|rhak|krol|jan|kag|ryr)(<vc>|in|or|an|ar|och|un|mar|yk|ja|arn|ir|ros|ror)" +
		"(|(mund|ard|arn|karr|chim|kos|rir|arl|kni|var|an|in|ir|a|i|as))"
	FANTASY_J_G_Z = "(aj|ch|etz|etzl|tz|kal|gahn|kab|aj|izl|ts|jaj|lan|kach|chaj|qaq|jol|ix|az|biq|nam)" +
		"(|(<vc>|aw|al|yes|il|ay|en|tom||oj|im|ol|aj|an|as))" +
		"(aj|am|al|aqa|ende|elja|ich|ak|ix|in|ak|al|il|ek|ij|os|al|im)"
	FANTASY_K_J_Y = "(yi|shu|a|be|na|chi|cha|cho|ksa|yi|shu)" +
		"(th|dd|jj|sh|rr|mk|n|rk|y|jj|th)" +
		"(us|ash|eni|akra|nai|ral|ect|are|el|urru|aja|al|uz|ict|arja|ichi|ural|iru|aki|esh)"
	FANTASY_S_E = "(syth|sith|srr|sen|yth|ssen|then|fen|ssth|kel|syn|est|bess|inth|nen|tin|cor|" +
		"sv|iss|ith|sen|slar|ssil|sthen|svis|s|ss|s|ss)" +
		"(|(tys|eus|yn|of|es|en|ath|elth|al|ell|ka|ith|yrrl|is|isl|yr|ast|iy))" +
		"(us|yn|en|ens|ra|rg|le|en|ith|ast|zon|in|yn|ys)"
)

type Template struct {
	Name string
	Tmpl string
}

var Templates = []Template{
	{Name: "MiddleEarth", Tmpl: MIDDLE_EARTH},
	{Name: "JapaneseNamesConstrained", Tmpl: JAPANESE_NAMES_CONSTRAINED},
	{Name: "JapaneseNamesDiverse", Tmpl: JAPANESE_NAMES_DIVERSE},
	{Name: "ChineseNames", Tmpl: CHINESE_NAMES},
	{Name: "GreekNames", Tmpl: GREEK_NAMES},
	{Name: "HawaiianNames1", Tmpl: HAWAIIAN_NAMES_1},
	{Name: "HawaiianNames2", Tmpl: HAWAIIAN_NAMES_2},
	{Name: "OldLatinPlaceNames", Tmpl: OLD_LATIN_PLACE_NAMES},
	{Name: "DragonsPern", Tmpl: DRAGONS_PERN},
	{Name: "DragonRiders", Tmpl: DRAGON_RIDERS},
	{Name: "Pokemon", Tmpl: POKEMON},
	{Name: "FantasyR", Tmpl: FANTASY_VOWELS_R},
	{Name: "FantasySA", Tmpl: FANTASY_S_A},
	{Name: "FantasyHL", Tmpl: FANTASY_H_L},
	{Name: "FantasyNL", Tmpl: FANTASY_N_L},
	{Name: "FantasyKN", Tmpl: FANTASY_K_N},
	{Name: "FantasyJGZ", Tmpl: FANTASY_J_G_Z},
	{Name: "FantasyKJY", Tmpl: FANTASY_K_J_Y},
	{Name: "FantasySE", Tmpl: FANTASY_S_E},
	{Name: "Funny", Tmpl: "sdD"},
	{Name: "Idiots", Tmpl: "ii"},
}

const (
	VERSION string = "0.0.1"
	Usage   string = `This is gfn, a fantasy name generator cli.

Usage: gfn [-vlc] <name|code>

Options:
-c --count    How many names to generate
-l --list     List pre-compiled shortcuts
-v --version  Show program version

pattern  syntax The  letters s,  v, V,  c, B,  C, i,  m, M,  D, and  d
represent different types of random replacements:

s - generic syllable
v - vowel
V - vowel or vowel combination
c - consonant
B - consonant or consonant combination suitable for beginning a word
C - consonant or consonant combination suitable anywhere in a word
i - insult
m - mushy name
M - mushy name ending
D - consonant suited for a stupid person's name
d - syllable suited for a stupid person's name (begins with a vowel)
Everything else is emitted literally.

All  characters  between parenthesis  ()  are  emitted literally.  For
example, the pattern s(dim), emits  a random generic syllable followed
by dim.

Characters  between angle  brackets <>  emit patterns  from the  table
above. Imagine the entire pattern is wrapped in one of these.

In  both  types of  groupings,  a  vertical  bar  | denotes  a  random
choice. Empty groups are allowed.  For example, (foo|bar) emits either
foo or bar. The pattern <c|v|>  emits a constant, vowel, or nothing at
all.

An exclamation point ! means  to capitalize the component that follows
it. For  example, !(foo) will emit  Foo and v!s will  emit a lowercase
vowel followed by a capitalized syllable, like eRod.`
)

type Config struct {
	Showversion   bool   `koanf:"version"` // -v
	Listshortcuts bool   `koanf:"list"`    // -l
	Code          string // arg
}

func InitConfig(output io.Writer) (*Config, error) {
	var kloader = koanf.New(".")

	// Load default values using the confmap provider.
	if err := kloader.Load(confmap.Provider(map[string]interface{}{
		"count": 10,
	}, "."), nil); err != nil {
		return nil, fmt.Errorf("failed to load default values into koanf: %w", err)
	}

	// setup custom usage
	flagset := flag.NewFlagSet("config", flag.ContinueOnError)
	flagset.Usage = func() {
		fmt.Fprintln(output, Usage)
		os.Exit(0)
	}

	// parse commandline flags
	flagset.BoolP("list", "l", false, "show list of precompiled codes")
	flagset.BoolP("version", "V", false, "show program version")

	if err := flagset.Parse(os.Args[1:]); err != nil {
		return nil, fmt.Errorf("failed to parse program arguments: %w", err)
	}

	// command line setup
	if err := kloader.Load(posflag.Provider(flagset, ".", kloader), nil); err != nil {
		return nil, fmt.Errorf("error loading flags: %w", err)
	}

	// fetch values
	conf := &Config{}
	if err := kloader.Unmarshal("", &conf); err != nil {
		return nil, fmt.Errorf("error unmarshalling: %w", err)
	}

	// arg is the code
	if len(flagset.Args()) > 0 {
		conf.Code = flagset.Args()[0]
	}

	return conf, nil
}
