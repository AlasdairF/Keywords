package keywords

import (
 "github.com/AlasdairF/Tokenize"
 "github.com/AlasdairF/Deaccent"
 "sort"
)

var stopwords = map[string]bool{
	`abbiamo`:false, `aber`:false, `able`:false, `about`:false, `abschnitt`:false, `aca`:false, `acerca`:false, `acht`:false, `achtzig`:false, `adelante`:false, `adeo`:false, `adesso`:false, `agli`:false, `ahora`:false, `ainsi`:false, `algun`:false,
	`alguna`:false, `algunas`:false, `alguno`:false, `algunos`:false, `aliquot`:false, `all`:false, `alla`:false, `alle`:false, `aller`:false, `alles`:false, `alli`:false, `allo`:false, `allora`:false, `alors`:false, `als`:false, `also`:false,
	`altre`:false, `altri`:false, `altro`:false, `ambos`:false, `ampleamos`:false, `anche`:false, `ancora`:false, `and`:false, `andare`:false, `andato`:false, `ante`:false, `antes`:false, `any`:false, `aparente`:false, `apart`:false, `aparte`:false,
	`apparente`:false, `apres`:false, `apud`:false, `aquel`:false, `aquella`:false, `aquellas`:false, `aquellos`:false, `aqui`:false, `arbitro`:false, `are`:false, `arent`:false, `arriba`:false, `asi`:false, `assez`:false, `atque`:false, `atras`:false,
	`auch`:false, `aucun`:false, `aucuns`:false, `auf`:false, `aun`:false, `aunque`:false, `aura`:false, `aurait`:false, `aurez`:false, `aus`:false, `auseinander`:false, `aussi`:false, `aut`:false, `autre`:false, `autres`:false, `aux`:false,
	`avaient`:false, `avait`:false, `avant`:false, `avec`:false, `avere`:false, `aveva`:false, `avevano`:false, `avez`:false, `avoir`:false, `avons`:false, `avrai`:false, `avrebbe`:false, `ayant`:false, `bajo`:false, `bald`:false, `bastante`:false,
	`beaucoup`:false, `become`:false, `becomes`:false, `becoming`:false, `been`:false, `bei`:false, `beide`:false, `beiden`:false, `being`:false, `ben`:false, `bene`:false, `ber`:false, `bien`:false, `bientot`:false, `bin`:false, `bis`:false,
	`bist`:false, `boite`:false, `bon`:false, `both`:false, `bueno`:false, `buono`:false, `but`:false, `cache`:false, `cada`:false, `came`:false, `camino`:false, `can`:false, `cannot`:false, `cant`:false, `capable`:false, `capannone`:false, `capaz`:false,
	`car`:false, `cascara`:false, `ce`:false, `cela`:false, `celle`:false, `celui`:false, `cependant`:false, `cero`:false, `ces`:false, `cest`:false, `cet`:false, `cette`:false, `ceux`:false, `chaque`:false, `che`:false, `chez`:false, `chi`:false,
	`ciao`:false, `cierta`:false, `ciertas`:false, `cierto`:false, `ciertos`:false, `cinco`:false, `cinq`:false, `cinque`:false, `cinquieme`:false, `circa`:false, `cobertizo`:false, `com`:false, `come`:false, `comes`:false, `comme`:false, `comment`:false,
	`como`:false, `comos`:false, `comprare`:false, `con`:false, `consecutivi`:false, `consecutivo`:false, `conseguimos`:false, `conseguir`:false, `consigo`:false, `consigue`:false, `consiguen`:false, `consigues`:false, `convertirse`:false, `convierte`:false,
	`coquille`:false, `cosa`:false, `cosi`:false, `costumbre`:false, `could`:false, `couldnt`:false, `coutume`:false, `cual`:false, `cualquier`:false, `cualquiera`:false, `cuando`:false, `cuatro`:false, `cui`:false, `cujus`:false, `cum`:false, `dabei`:false,
	`dadurch`:false, `daher`:false, `dal`:false, `dalla`:false, `damit`:false, `dann`:false, `dans`:false, `darum`:false, `das`:false, `dass`:false, `dautres`:false, `dazu`:false, `debe`:false, `debera`:false, `deberia`:false, `debido`:false, `debut`:false,
	`decir`:false, `dedans`:false, `dehors`:false, `dei`:false, `dein`:false, `deinde`:false, `deine`:false, `dejar`:false, `del`:false, `dell`:false, `della`:false, `delle`:false, `dello`:false, `dem`:false, `demasiado`:false, `den`:false, `denen`:false,
	`denn`:false, `dentro`:false, `depuis`:false, `der`:false, `deren`:false, `des`:false, `desde`:false,`deshalb`:false, `desiderare`:false, `dessen`:false, `detto`:false, `deux`:false, `deuxieme`:false, `deuxiemementzweitens`:false, `devait`:false,
	`devant`:false, `deve`:false, `devenir`:false, `devient`:false, `devo`:false, `devrait`:false, `dice`:false, `dicho`:false, `did`:false, `didnt`:false, `die`:false, `diejenigen`:false, `dies`:false, `diese`:false, `diesem`:false, `diesen`:false,
	`dieser`:false, `dieses`:false, `digitized`:false, `dire`:false, `dit`:false, `diventa`:false, `diventando`:false, `diventare`:false, `dix`:false, `doch`:false, `does`:false, `doesnt`:false, `doing`:false, `doit`:false, `donc`:false, `donde`:false,
	`done`:false, `donne`:false, `donner`:false, `dont`:false, `doppio`:false, `dort`:false, `dos`:false, `dove`:false, `dovere`:false, `dovra`:false, `dovrebbe`:false, `dovuto`:false, `drei`:false, `dritte`:false, `droite`:false, `due`:false, `dum`:false,
	`dun`:false, `dune`:false, `durante`:false, `durch`:false, `during`:false, `each`:false, `eam`:false, `eben`:false, `ecco`:false, `egli`:false, `eher`:false, `eight`:false, `eighty`:false, `ein`:false, `eine`:false, `einem`:false, `einen`:false,
	`einer`:false, `eines`:false, `einnahme`:false, `either`:false, `ejus`:false, `ella`:false, `ellas`:false, `elle`:false, `elles`:false, `ellos`:false, `else`:false, `empleais`:false, `emplean`:false, `emplear`:false, `empleas`:false, `empleo`:false,
	`encima`:false, `encore`:false, `entonces`:false, `entrambi`:false, `entre`:false, `entweder`:false, `era`:false, `eramos`:false, `eran`:false, `erano`:false, `erant`:false, `eras`:false, `erat`:false, `eres`:false, `erst`:false, `erste`:false,
	`ersten`:false, `esempio`:false, `eso`:false, `essai`:false, `esse`:false, `essere`:false, `esset`:false, `essi`:false, `esso`:false, `est`:false, `esta`:false, `estaba`:false, `estado`:false, `estais`:false, `estamos`:false, `estan`:false,
	`este`:false, `esto`:false, `estos`:false, `estoy`:false, `etaient`:false, `etait`:false, `etant`:false, `etat`:false, `etc`:false, `ete`:false, `etes`:false, `etiam`:false, `etions`:false, `etre`:false, `euer`:false, `eum`:false, `eure`:false,
	`eut`:false, `eux`:false, `even`:false, `ever`:false, `every`:false, `facendo`:false, `fahig`:false, `faire`:false, `faisait`:false, `fait`:false, `faites`:false, `fallig`:false, `faranno`:false, `fare`:false, `faresti`:false, `faro`:false,
	`fatto`:false, `faut`:false, `fere`:false, `feriez`:false, `few`:false, `fifth`:false, `fin`:false, `fine`:false, `fino`:false, `first`:false, `fit`:false, `five`:false, `fois`:false, `font`:false, `for`:false, `force`:false, `forth`:false,
	`four`:false, `fra`:false, `from`:false, `fue`:false, `fueron`:false, `fui`:false, `fuimos`:false, `fuit`:false, `funf`:false, `funfte`:false, `fur`:false, `furent`:false, `fut`:false, `gabe`:false, `ganz`:false, `gar`:false, `gehen`:false,
	`geht`:false, `genommen`:false, `gente`:false, `gerade`:false, `gergo`:false, `gewesen`:false, `gia`:false, `gibt`:false, `giebt`:false, `gilt`:false, `ging`:false, `giu`:false, `gleich`:false, `gli`:false, `goes`:false, `grado`:false, `gueno`:false,
	`guscio`:false, `habe`:false, `haben`:false, `haber`:false, `habia`:false, `habian`:false, `habra`:false, `habria`:false, `hace`:false, `haceis`:false, `hacemos`:false, `hacen`:false, `hacer`:false, `haces`:false, `haciendo`:false, `had`:false,
	`hadnt`:false, `haec`:false, `hago`:false, `hai`:false, `hallo`:false, `han`:false, `hanc`:false, `hangar`:false, `hanno`:false, `haria`:false, `has`:false, `hasnt`:false, `hasta`:false, `hat`:false, `hatte`:false, `hatten`:false, `hattest`:false,
	`hattet`:false, `haut`:false, `have`:false, `havent`:false, `having`:false, `hay`:false, `hecho`:false, `hed`:false, `heisst`:false, `hemos`:false, `her`:false, `here`:false, `heres`:false, `hes`:false, `hid`:false, `hier`:false, `hierher`:false,
	`him`:false, `hinter`:false, `hipocresia`:false, `his`:false, `hither`:false, `hizo`:false, `hoc`:false, `hola`:false, `hors`:false, `how`:false, `hows`:false, `huit`:false, `hunc`:false, `iba`:false, `ich`:false, `ici`:false, `idem`:false, `ihm`:false,
	`ihn`:false, `ihnen`:false, `ihr`:false, `ihre`:false, `ihrer`:false, `iii`:false, `iis`:false, `ill`:false, `ille`:false, `illi`:false, `illis`:false, `illum`:false, `ils`:false, `immer`:false, `inc`:false, `incluso`:false, `inde`:false, `indem`:false,
	`index`:false, `indice`:false, `indietro`:false, `ins`:false, `intenta`:false, `intentais`:false, `intentamos`:false, `intentan`:false, `intentar`:false, `intentas`:false, `intento`:false, `inter`:false, `into`:false, `invece`:false, `ipse`:false,
	`isnt`:false, `ist`:false, `ita`:false, `itd`:false, `itll`:false, `its`:false, `ive`:false, `jai`:false, `jam`:false, `jamais`:false, `je`:false, `jede`:false, `jedem`:false, `jeden`:false, `jeder`:false, `jedes`:false, `jene`:false, `jener`:false,
	`jenes`:false, `jetzt`:false, `jusqua`:false, `just`:false, `juste`:false, `kam`:false, `kann`:false, `kannst`:false, `kein`:false, `kippen`:false, `knnen`:false, `kommen`:false, `kommt`:false, `konnen`:false, `konnt`:false, `konnte`:false,
	`laisser`:false, `largo`:false, `las`:false, `lascia`:false, `lasciare`:false, `lassen`:false, `lata`:false, `lattina`:false, `lautre`:false, `lavoro`:false, `lei`:false, `les`:false, `lest`:false, `let`:false, `lets`:false, `leur`:false, `leurs`:false,
	`loro`:false, `los`:false, `lugar`:false, `lui`:false, `lungo`:false, `luogo`:false, `machen`:false, `macht`:false, `mai`:false, `maintenant`:false, `mais`:false, `man`:false, `maniere`:false, `marier`:false, `mas`:false, `may`:false, `meglio`:false,
	`mehr`:false, `mein`:false, `meine`:false, `meme`:false, `ment`:false, `mentre`:false, `mes`:false, `mettere`:false, `mettre`:false, `mientras`:false, `mine`:false, `mio`:false, `mir`:false, `mismo`:false, `mit`:false, `modo`:false, `moi`:false,
	`moins`:false, `molta`:false, `molti`:false, `molto`:false, `mon`:false, `mosto`:false, `mot`:false, `muchos`:false, `muss`:false, `mussen`:false, `musst`:false, `must`:false, `mustnt`:false, `muy`:false, `nach`:false, `nachdem`:false, `nec`:false,
	`necesario`:false, `nehmen`:false, `nei`:false, `neigt`:false, `nein`:false, `nel`:false, `nella`:false, `nelle`:false, `neque`:false, `nest`:false, `netaient`:false, `netait`:false, `neuf`:false, `neun`:false, `neunzig`:false, `nicht`:false,
	`nine`:false, `ninety`:false, `noch`:false, `noi`:false, `nome`:false, `nommes`:false, `non`:false, `nont`:false, `nor`:false, `nos`:false, `nosotros`:false, `nostro`:false, `not`:false, `notre`:false, `nous`:false, `nouveaux`:false, `novanta`:false,
	`nove`:false, `noventa`:false, `now`:false, `nueve`:false, `null`:false, `nun`:false, `nunca`:false, `nuovi`:false, `nuovo`:false, `nur`:false, `oben`:false, `obwohl`:false, `ochenta`:false, `ocho`:false, `oder`:false, `ogni`:false, `ohne`:false,
	`oltre`:false, `omnia`:false, `one`:false, `ones`:false, `only`:false, `ont`:false, `onto`:false, `opiz`:false, `ora`:false, `otro`:false, `ottanta`:false, `otto`:false, `ought`:false, `our`:false, `par`:false, `para`:false, `paraitre`:false,
	`parce`:false, `parece`:false, `parecer`:false, `parecido`:false, `parole`:false, `part`:false, `parte`:false, `partir`:false, `pas`:false, `paulo`:false, `paura`:false, `peggio`:false, `pendant`:false, `pente`:false, `per`:false, `permet`:false,
	`pero`:false, `persone`:false, `personnes`:false, `peu`:false, `peur`:false, `peut`:false, `piacerebbe`:false, `piece`:false, `piu`:false, `piuttosto`:false, `plupart`:false, `plus`:false, `plusieurs`:false, `plutot`:false, `pochi`:false, `poco`:false,
	`pocos`:false, `podeis`:false, `podemos`:false, `poder`:false, `podia`:false, `podra`:false, `podria`:false, `podriais`:false, `podriamos`:false, `podrian`:false, `podrias`:false, `poi`:false, `poner`:false, `por`:false, `porque`:false, `post`:false,
	`poteva`:false, `potuto`:false, `pour`:false, `pourquoi`:false, `pourrait`:false, `pouvait`:false, `prae`:false, `premiere`:false, `prendere`:false, `prendre`:false, `pres`:false, `presa`:false, `preso`:false, `presto`:false, `primero`:false,
	`primo`:false, `pris`:false, `prise`:false, `pro`:false, `promesso`:false, `pronto`:false, `puede`:false, `pueden`:false, `puedo`:false, `pues`:false, `puis`:false, `puo`:false, `put`:false, `qua`:false, `quae`:false, `quale`:false, `qualsiasi`:false,
	`quam`:false, `quand`:false, `quando`:false, `quanto`:false, `quarto`:false, `quasi`:false, `quatre`:false, `quattro`:false, `que`:false, `quel`:false, `quella`:false, `quelle`:false, `quelles`:false, `quelli`:false, `quello`:false, `quelques`:false,
	`quels`:false, `quem`:false, `querer`:false, `questa`:false, `queste`:false, `questo`:false, `qui`:false, `quibus`:false, `quid`:false, `quidem`:false, `quien`:false, `quiere`:false, `quil`:false, `quils`:false, `quin`:false, `quindi`:false,
	`quinto`:false, `quo`:false, `quod`:false, `quon`:false, `quoque`:false, `quorum`:false, `quos`:false, `quun`:false, `raccontare`:false, `ran`:false, `rather`:false, `ref`:false, `rem`:false, `rispetto`:false, `rum`:false, `sabe`:false, `sabeis`:false,
	`sabemos`:false, `saben`:false, `saber`:false, `sabes`:false, `sage`:false, `sagen`:false, `sagt`:false, `sagte`:false, `said`:false, `salut`:false, `same`:false, `sans`:false, `sara`:false, `sarebbe`:false, `saw`:false, `say`:false, `saying`:false,
	`says`:false, `schale`:false, `scheinbaren`:false, `scheinen`:false, `scheint`:false, `schien`:false, `schuppen`:false, `scie`:false, `sec`:false, `seccion`:false, `sechs`:false, `second`:false, `seconde`:false, `secondly`:false, `secondo`:false,
	`section`:false, `sed`:false, `seem`:false, `seemed`:false, `seeming`:false, `seems`:false, `sega`:false, `segundo`:false, `seguro`:false, `sehr`:false, `sei`:false, `seid`:false, `sein`:false, `seine`:false, `seinem`:false, `seinen`:false,
	`seiner`:false, `seines`:false, `seis`:false, `sekunde`:false, `selbst`:false, `semblait`:false, `semble`:false, `sembra`:false, `sembrare`:false, `sembrava`:false, `senza`:false, `sept`:false, `ser`:false, `serais`:false, `serait`:false,
	`seria`:false, `ses`:false, `sette`:false, `setzen`:false, `seule`:false, `seulement`:false, `seven`:false, `sezione`:false, `shall`:false, `shant`:false, `she`:false, `shed`:false, `shell`:false, `shes`:false, `should`:false, `shouldnt`:false,
	`sia`:false, `siamo`:false, `sibi`:false, `sic`:false, `sich`:false, `sicher`:false, `sicuro`:false, `sie`:false, `sieben`:false, `sien`:false, `siendo`:false, `sierra`:false, `siete`:false, `sin`:false, `sind`:false, `sino`:false, `six`:false,
	`sobre`:false, `sois`:false, `soit`:false, `solamente`:false, `soll`:false, `sollen`:false, `sollst`:false, `sollt`:false, `sollte`:false, `solo`:false, `somos`:false, `son`:false, `sondern`:false, `sono`:false, `sonst`:false, `sont`:false,
	`soon`:false, `sopra`:false, `soprattutto`:false, `sotto`:false, `sous`:false, `souvent`:false, `soweit`:false, `sowie`:false, `soy`:false, `soyez`:false, `sposarsi`:false, `spruch`:false, `stati`:false, `stato`:false, `stesso`:false, `still`:false,
	`sua`:false, `sub`:false, `subito`:false, `such`:false, `suddetto`:false, `sue`:false, `suis`:false, `sujet`:false, `sul`:false, `sulla`:false, `sunt`:false, `suo`:false, `suoi`:false, `sup`:false, `sur`:false, `sure`:false, `sus`:false, `suyos`:false,
	`taient`:false, `tait`:false, `take`:false, `taken`:false, `taking`:false, `tal`:false, `tale`:false, `tam`:false, `tambien`:false, `tamen`:false, `tan`:false, `tandis`:false, `tant`:false, `tanto`:false, `tat`:false, `tel`:false, `tell`:false,
	`tellement`:false, `tels`:false, `tempo`:false, `tend`:false, `tende`:false, `tends`:false, `teneis`:false, `tenemos`:false, `tener`:false, `tengo`:false, `tenia`:false, `tenido`:false, `tercera`:false, `terzo`:false, `tes`:false, `than`:false,
	`that`:false, `thatll`:false, `thats`:false, `thatve`:false, `the`:false, `their`:false, `theirs`:false, `them`:false, `then`:false, `thence`:false, `there`:false, `thered`:false, `therell`:false, `therere`:false, `theres`:false, `thereve`:false,
	`these`:false, `they`:false, `theyd`:false, `theyll`:false, `theyre`:false, `theyve`:false, `third`:false, `this`:false, `those`:false, `thou`:false, `though`:false, `three`:false, `thru`:false, `thun`:false, `thus`:false, `tiempo`:false,
	`tiende`:false, `tiene`:false, `tienen`:false, `tienes`:false, `til`:false, `todavia`:false, `todo`:false, `toma`:false, `tomado`:false, `tomar`:false, `ton`:false, `too`:false, `toujours`:false, `tous`:false, `tout`:false, `toute`:false,
	`toutes`:false, `tra`:false, `trabaja`:false, `trabajais`:false, `trabajamos`:false, `trabajan`:false, `trabajar`:false, `trabajas`:false, `trabajo`:false, `tras`:false, `travers`:false, `tre`:false, `tres`:false, `triplo`:false, `trois`:false,
	`troisieme`:false, `trop`:false, `troppo`:false, `trouve`:false, `trova`:false, `tum`:false, `tun`:false, `tur`:false, `tut`:false, `tutto`:false, `tuyo`:false, `twice`:false, `two`:false, `uber`:false, `ubi`:false, `uhr`:false, `ultimo`:false,
	`una`:false, `unas`:false, `und`:false, `une`:false, `uno`:false, `unos`:false, `uns`:false, `unser`:false, `unsere`:false, `unter`:false, `unto`:false, `upon`:false, `ups`:false, `usa`:false, `usais`:false, `usamos`:false, `usan`:false,
	`usar`:false, `usas`:false, `uso`:false, `usted`:false, `vai`:false, `vais`:false, `valeur`:false, `valor`:false, `vamos`:false, `van`:false, `vas`:false, `vaya`:false, `veces`:false, `vel`:false, `venir`:false, `venire`:false, `venu`:false,
	`venuto`:false, `ver`:false, `verdad`:false, `verdadera`:false, `verdadero`:false, `vero`:false, `vers`:false, `verso`:false, `versteckte`:false, `very`:false, `veut`:false, `via`:false, `vida`:false, `viene`:false, `vient`:false, `vier`:false,
	`vingt`:false, `vingts`:false, `vino`:false, `voi`:false, `voici`:false, `voie`:false, `voient`:false, `voila`:false, `volonta`:false, `volonte`:false, `volte`:false, `voluntad`:false, `vom`:false, `von`:false, `vont`:false, `vor`:false,
	`vosotras`:false, `vosotros`:false, `vostro`:false, `votre`:false, `voudrais`:false, `vouloir`:false, `vous`:false, `voy`:false, `vuole`:false, `wahrend`:false, `wann`:false, `want`:false, `wants`:false, `war`:false, `ware`:false, `waren`:false,
	`warum`:false, `was`:false, `wasnt`:false, `way`:false, `wed`:false, `weg`:false, `weil`:false, `weise`:false, `weiter`:false, `weitere`:false, `welche`:false, `welcher`:false, `welches`:false, `well`:false, `wenige`:false, `wenn`:false, `went`:false,
	`wer`:false, `werde`:false, `werden`:false, `werdet`:false, `were`:false, `werent`:false, `weshalb`:false, `weve`:false, `what`:false, `whatll`:false, `whats`:false, `when`:false, `whens`:false, `where`:false, `wheres`:false, `whether`:false,
	`which`:false, `while`:false, `who`:false, `whod`:false, `wholl`:false, `whos`:false, `wie`:false, `wieder`:false, `wieso`:false, `will`:false, `wir`:false, `wird`:false, `wirst`:false, `with`:false, `woher`:false, `wohin`:false, `wol`:false,
	`wollen`:false, `wont`:false, `would`:false, `wouldnt`:false, `wre`:false, `wurde`:false, `wurden`:false, `wurdest`:false, `yet`:false, `you`:false, `youd`:false, `youll`:false, `your`:false, `yours`:false, `youve`:false, `zero`:false, `zugleich`:false,
	`zum`:false, `zur`:false, `zwar`:false, `zwei`:false, `zweimal`:false, `zweite`:false, `some`:false, `most`:false, `other`:false, `because`:false, `many`:false, `much`:false, 
}

type keyVal struct {
	k string
	v uint32
}
type sorter []keyVal
func (a sorter) Len() int           { return len(a) }
func (a sorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter) Less(i, j int) bool { return a[i].v > a[j].v } // descending order

type keywordsStruct struct {
	km map[string]uint32
}

func New() *keywordsStruct {
	k := new(keywordsStruct)
	k.km = make(map[string]uint32)
	return k
}

func (k *keywordsStruct) Monogram(b []byte) {
	var err error
	var d []byte
	var ok bool
	wordfn := func(word []byte) {
		if d, err = deaccent.Bytes(word); err == nil {
			if len(d) >= 3 {
				if _, ok = stopwords[string(d)]; !ok {
					k.km[string(word)]++
				}
			}
		}
	}
	tokenize.AllInOne(b, wordfn, true, false, true, true, false)
}

func (k *keywordsStruct) Bigram(b []byte) {
	var err error
	var d []byte
	var ok, last bool
	var lastword, word1 string
	wordfn := func(word []byte) {
		if d, err = deaccent.Bytes(word); err == nil {
			if len(d) >= 3 {
				if _, ok = stopwords[string(d)]; !ok {
					word1 = string(word)
					if last {
						k.km[lastword + ` ` + word1]++
					}
					lastword = word1
					last = true
				} else {
					last = false
				}
			} else {
				last = false
			}
		} else {
			last = false
		}
	}
	tokenize.AllInOne(b, wordfn, true, false, true, true, false)
}

func (k *keywordsStruct) Ngram(b []byte) {
	var err error
	var d []byte
	var ok, last bool
	var lastword, word1 string
	wordfn := func(word []byte) {
		if d, err = deaccent.Bytes(word); err == nil {
			if len(d) >= 3 {
				if _, ok = stopwords[string(d)]; !ok {
					word1 = string(word)
					if last {
						k.km[lastword + ` ` + word1]++
					} else {
						k.km[word1]++
					}
					lastword = word1
					last = true
				} else {
					last = false
				}
			} else {
				last = false
			}
		} else {
			last = false
		}
	}
	tokenize.AllInOne(b, wordfn, true, false, true, true, false)
}

func (k *keywordsStruct) Result(num int) []string {
	lst := make(sorter, len(k.km))
	var i int
	for word, v := range k.km {
		lst[i] = keyVal{word, v}
		i++
	}
	sort.Sort(lst)
	if num > i {
		num = i
	}
	res := make([]string, num)
	for i=0; i<num; i++ {
		res[i] = lst[i].k
	}
	return res
}
