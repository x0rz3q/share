package main

import (
	"hash/fnv"
	"math/rand"
	"strings"
)

var (
	left = [...]string{
		"admiring",
		"adoring",
		"affectionate",
		"agitated",
		"amazing",
		"angry",
		"awesome",
		"beautiful",
		"blissful",
		"bold",
		"boring",
		"brave",
		"busy",
		"charming",
		"clever",
		"cool",
		"compassionate",
		"competent",
		"condescending",
		"confident",
		"cranky",
		"crazy",
		"dazzling",
		"determined",
		"distracted",
		"dreamy",
		"eager",
		"ecstatic",
		"elastic",
		"elated",
		"elegant",
		"eloquent",
		"epic",
		"exciting",
		"fervent",
		"festive",
		"flamboyant",
		"focused",
		"friendly",
		"frosty",
		"funny",
		"gallant",
		"gifted",
		"goofy",
		"gracious",
		"great",
		"happy",
		"hardcore",
		"heuristic",
		"hopeful",
		"hungry",
		"infallible",
		"inspiring",
		"interesting",
		"intelligent",
		"jolly",
		"jovial",
		"keen",
		"kind",
		"laughing",
		"loving",
		"lucid",
		"magical",
		"mystifying",
		"modest",
		"musing",
		"naughty",
		"nervous",
		"nice",
		"nifty",
		"nostalgic",
		"objective",
		"optimistic",
		"peaceful",
		"pedantic",
		"pensive",
		"practical",
		"priceless",
		"quirky",
		"quizzical",
		"recursing",
		"relaxed",
		"reverent",
		"romantic",
		"sad",
		"serene",
		"sharp",
		"silly",
		"sleepy",
		"stoic",
		"strange",
		"stupefied",
		"suspicious",
		"sweet",
		"tender",
		"thirsty",
		"trusting",
		"unruffled",
		"upbeat",
		"vibrant",
		"vigilant",
		"vigorous",
		"wizardly",
		"wonderful",
		"xenodochial",
		"youthful",
		"zealous",
		"zen",
	}
	right = strings.Fields(`fawn peacock fox terrier civet musk deer seastar pigeon bull bumblebee crocodile flying squirrel elephant leopard seal baboon porcupine wolverine spider monkey vampire bat sparrow manatee possum swallow wildcat bandicoot labradoodle dragonfly tarsier snowy owl chameleon boykin puffin bison llama kitten stinkbug macaw parrot leopard cat prawn panther dogfish fennec frigatebird nurse shark turkey cockatoo neanderthal crow gopher reindeer earwig  anaconda panda ant silver fox collared peccary puppy common buzzard moose binturong wildebeest lovebird ferret persian marine toad woolly mammoth dalmatian bird umbrellabird kingfisher kangaroo stallion russian blue ostrich owl tawny owl affenpinscher caiman elephant seal octopus meerkat whale shark buck donkey red wolf mountain lion labrador retriever quetzal chamois sponge hamster orangutan sea urchin uakari doberman dormouse saint bernard bull shark ocelot sparrow spitz stoat snapping turtle dragonfly cougar alligator walrus glass lizard malayan tiger frog tiger armadillo chinchilla crab squid calf shrew dolphin royal penguin dingo turtle yellow-eyed penguin chimpanzee armadillo boa constrictor rabbit basking coyote chinook osprey sea lion fly sperm whale patas monkey tiffany mountain goat dodo worm cat warthog peccary shark pony monkey swan whippet beagle cougar anteater quail liger cheetah woodpecker egret eagle moose warthog honey bee snail stag beetle budgie molly magpie rhinoceros elephant kudu wombat tree frog goat lamb tropicbird human hog tang pool frog lemur ox dog lizard echidna great dane wallaby hawk dove jellyfish sloth macaque starfish sun bear guppy welsh corgi deer impala porpoise gazelle bichon seal wolf zebra shark mole narwhal hedgehog sheep horse bluetick colt spadefoot toad wildebeest piranha basenji mallard bull mastiff bear siberian husky bird badger red panda hammerhead rock hyrax kangaroo marsh frog mule weasel dogfish dachsbracke forest elephant oyster bat python coati platypus salamander cat caterpillar giraffe snake kid falcon robin guinea fowl tern sea lion dingo bolognese drake goose rat gentoo penguin iguana quail mouse horseshoe crab roebuck cattle dog fish poodle frog wolverine chinchilla bobcat grey seal hermit crab carolina shepherd gila monster snail mandrill leopard frilled lizard echidna rabbit bison barracuda foal ass eagle octopus avocet siamese dodo yorkie cockroach wallaroo tiger woodlouse glow worm fossa buffalo zorse albatross indri seahorse lemur louse ostrich humpback whale millipede fin whale joey pinscher dachshund proboscis monkey pelican chihuahua dogo indian rhinoceros wasp siberian raccoon dog yak stingray jack russel water vole foxhound sheep stork horse monkey woolly monkey waterbuck dunker cuscus ibis giraffe aardvark hummingbird grizzly bear otter pike minke whale pika stickbug pelican dugong bongo lemming shrimp piglet sabre-toothed tiger gemsbok tiger shark tuatara rottweiler elephant shrew ewe coati cichlid akita gharial thorny devil duck macaroni penguin steer setter pufferfish donkey mink macaw wolfhound white tiger ram ant rat marten galapagos tortoise crab horn shark blue whale koala starfish partridge sea squirt fire-bellied toad chipmunk ibex maltese clumber butterfly manta ray flamingo opossum parrot mastiff water buffalo okapi salmon tapir adelie killer whale lynx basilisk indian elephant oyster manta ray prairie dog chipmunk locust dog cottontop hyena spectacled bear oriole cobra pug monitor mandrill antelope chinstrap zebra chicken mule seal goat little penguin gull tasmanian devil caterpillar tamarin wrasse woodchuck otter penguin porcupine killer whale bear ferret dusky nightingale slow worm bat jaguar humboldt ermine saola emu lobster weasel nightingale hound bombay platypus electric eel asian elephant sea otter uguisu scorpion fox jerboa bengal tiger zebu lion zonkey ragdoll caracal bee kiwi puma common loon jackal malamute mayfly baboon terrier jellyfish vicuna penguin desert tortoise muskrat water dragon zebra malayan civet burmese orangutan himalayan pond skater howler monkey newt border collie cow bearded dragon fish barn owl puffin chin anteater beaver canary hamster sloth collie heron sea dragon gopher magpie king crab flounder opossum pademelon capybara boar leaf-tailed gecko turkey clown fish musk-ox bulldog pronghorn hercules beetle reindeer llama pygmy eskimo dog kinkajou komodo dragon cuttlefish cub bloodhound squirrel gander moorhen emu brown bear javanese birman harrier tortoise antelope gnu kingfisher wasp olm havanese canaan lizard indochinese tiger ocelot mist hare discus cony orca rooster ground hog silver dollar peacock akbash somali beaver maine coon mouse eland squirrel serval chimpanzee snowshoe toucan catfish lynx coyote bunny retriever fur seal cow balinese vulture coral leopard raccoon polar bear okapi kakapo whale sand lizard bonobo moray gila monster cormorant bracke camel markhor rockhopper neapolitan black bear roseate spoonbill woodpecker mountain lion crested penguin hippopotamus puma camel alligator guinea pig heron siberian tiger river dolphin axolotl argentino human mongoose drever quokka common frog elk wombat spider monkey civet sting ray panther gar lionfish snake crane newt raven tortoise fire ant chicadee common toad pig manatee centipede numbat river turtle falcon angelfish chamois rhinoceros shark flamingo pheasant ladybird grasshopper greyhound lemming pig marmoset eel yorkiepoo tiger salamander mosquito shih tzu quoll chick guanaco walrus badger ainu squid pekingese gerbil duck rattlesnake tapir lobster catfish mustang wallaby mongrel butterfly booby bush elephant fox rattlesnake cockroach tadpole lark ape pied tamarin mare tetra squirrel monkey elephant seal dhole cesky raccoon newfoundland marmoset stag bullfrog black bear crocodile lion barb wolf vervet monkey beetle polar bear pointer grizzly bear meerkat owl reptile fousek gibbon king penguin budgerigar swan hartebeest cassowary borneo elephant oryx alpaca gerbil chameleon galapagos penguin vulture barracuda insect fishing cat hen giant clam hare polecat fly chow chow yak seahorse spider eel burro brown bear boxer dog crane bandicoot hedgehog dromedary goose budgerigar dolphin kelpie dog highland cattle dormouse duckbill springbok mongoose bobcat water buffalo gecko hornet iguana wild boar koala guinea pig marmot skink deer filly barnacle tree toad leopard tortoise appenzeller doe gecko mole mynah bird mau gilla monster french bulldog termite salamander parakeet finch horned frog hippopotamus hummingbird cheetah albatross jaguar toad hyena gorilla skunk impala sea slug scorpion fish jackal skunk grouse sea turtle moth caribou dugong bighorn sheep ibizan hound gorilla puffer fish chicken komodo dragon buffalo`)
)

func RandomName(seedString string) string {
	h := fnv.New32a()
	h.Write([]byte(seedString))
	seed := int64(h.Sum32())
	src := rand.New(rand.NewSource(seed))

begin:
	l := left[src.Intn(len(left))]
	r := right[src.Intn(len(right))]
	if string(l[0]) != string(r[0]) {
		goto begin
	}
	return l + r
}
