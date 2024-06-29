package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

const sampleText = `The quick brown fox jumps over the lazy dog.`

var (
	jaKanjis = []rune{}
)

func init() {
	// table is the list of Japanese Kanji characters in a part of JIS X 0208.
	const table = `
亜唖娃阿哀愛挨姶逢葵茜穐悪握渥旭葦芦鯵梓圧斡扱宛姐虻飴絢綾鮎或粟袷安庵按暗案闇鞍杏以伊位依偉囲夷委威尉惟意慰易椅為畏異移維緯胃萎衣謂違遺医井亥域育郁磯一壱溢逸稲茨芋鰯允印咽員因姻引飲淫胤蔭
院陰隠韻吋右宇烏羽迂雨卯鵜窺丑碓臼渦嘘唄欝蔚鰻姥厩浦瓜閏噂云運雲荏餌叡営嬰影映曳栄永泳洩瑛盈穎頴英衛詠鋭液疫益駅悦謁越閲榎厭円園堰奄宴延怨掩援沿演炎焔煙燕猿縁艶苑薗遠鉛鴛塩於汚甥凹央奥往応
押旺横欧殴王翁襖鴬鴎黄岡沖荻億屋憶臆桶牡乙俺卸恩温穏音下化仮何伽価佳加可嘉夏嫁家寡科暇果架歌河火珂禍禾稼箇花苛茄荷華菓蝦課嘩貨迦過霞蚊俄峨我牙画臥芽蛾賀雅餓駕介会解回塊壊廻快怪悔恢懐戒拐改
魁晦械海灰界皆絵芥蟹開階貝凱劾外咳害崖慨概涯碍蓋街該鎧骸浬馨蛙垣柿蛎鈎劃嚇各廓拡撹格核殻獲確穫覚角赫較郭閣隔革学岳楽額顎掛笠樫橿梶鰍潟割喝恰括活渇滑葛褐轄且鰹叶椛樺鞄株兜竃蒲釜鎌噛鴨栢茅萱
粥刈苅瓦乾侃冠寒刊勘勧巻喚堪姦完官寛干幹患感慣憾換敢柑桓棺款歓汗漢澗潅環甘監看竿管簡緩缶翰肝艦莞観諌貫還鑑間閑関陥韓館舘丸含岸巌玩癌眼岩翫贋雁頑顔願企伎危喜器基奇嬉寄岐希幾忌揮机旗既期棋棄
機帰毅気汽畿祈季稀紀徽規記貴起軌輝飢騎鬼亀偽儀妓宜戯技擬欺犠疑祇義蟻誼議掬菊鞠吉吃喫桔橘詰砧杵黍却客脚虐逆丘久仇休及吸宮弓急救朽求汲泣灸球究窮笈級糾給旧牛去居巨拒拠挙渠虚許距鋸漁禦魚亨享京
供侠僑兇競共凶協匡卿叫喬境峡強彊怯恐恭挟教橋況狂狭矯胸脅興蕎郷鏡響饗驚仰凝尭暁業局曲極玉桐粁僅勤均巾錦斤欣欽琴禁禽筋緊芹菌衿襟謹近金吟銀九倶句区狗玖矩苦躯駆駈駒具愚虞喰空偶寓遇隅串櫛釧屑屈
掘窟沓靴轡窪熊隈粂栗繰桑鍬勲君薫訓群軍郡卦袈祁係傾刑兄啓圭珪型契形径恵慶慧憩掲携敬景桂渓畦稽系経継繋罫茎荊蛍計詣警軽頚鶏芸迎鯨劇戟撃激隙桁傑欠決潔穴結血訣月件倹倦健兼券剣喧圏堅嫌建憲懸拳捲
検権牽犬献研硯絹県肩見謙賢軒遣鍵険顕験鹸元原厳幻弦減源玄現絃舷言諺限乎個古呼固姑孤己庫弧戸故枯湖狐糊袴股胡菰虎誇跨鈷雇顧鼓五互伍午呉吾娯後御悟梧檎瑚碁語誤護醐乞鯉交佼侯候倖光公功効勾厚口向
后喉坑垢好孔孝宏工巧巷幸広庚康弘恒慌抗拘控攻昂晃更杭校梗構江洪浩港溝甲皇硬稿糠紅紘絞綱耕考肯肱腔膏航荒行衡講貢購郊酵鉱砿鋼閤降項香高鴻剛劫号合壕拷濠豪轟麹克刻告国穀酷鵠黒獄漉腰甑忽惚骨狛込
此頃今困坤墾婚恨懇昏昆根梱混痕紺艮魂些佐叉唆嵯左差査沙瑳砂詐鎖裟坐座挫債催再最哉塞妻宰彩才採栽歳済災采犀砕砦祭斎細菜裁載際剤在材罪財冴坂阪堺榊肴咲崎埼碕鷺作削咋搾昨朔柵窄策索錯桜鮭笹匙冊刷
察拶撮擦札殺薩雑皐鯖捌錆鮫皿晒三傘参山惨撒散桟燦珊産算纂蚕讃賛酸餐斬暫残仕仔伺使刺司史嗣四士始姉姿子屍市師志思指支孜斯施旨枝止死氏獅祉私糸紙紫肢脂至視詞詩試誌諮資賜雌飼歯事似侍児字寺慈持時
次滋治爾璽痔磁示而耳自蒔辞汐鹿式識鴫竺軸宍雫七叱執失嫉室悉湿漆疾質実蔀篠偲柴芝屡蕊縞舎写射捨赦斜煮社紗者謝車遮蛇邪借勺尺杓灼爵酌釈錫若寂弱惹主取守手朱殊狩珠種腫趣酒首儒受呪寿授樹綬需囚収周
宗就州修愁拾洲秀秋終繍習臭舟蒐衆襲讐蹴輯週酋酬集醜什住充十従戎柔汁渋獣縦重銃叔夙宿淑祝縮粛塾熟出術述俊峻春瞬竣舜駿准循旬楯殉淳準潤盾純巡遵醇順処初所暑曙渚庶緒署書薯藷諸助叙女序徐恕鋤除傷償
勝匠升召哨商唱嘗奨妾娼宵将小少尚庄床廠彰承抄招掌捷昇昌昭晶松梢樟樵沼消渉湘焼焦照症省硝礁祥称章笑粧紹肖菖蒋蕉衝裳訟証詔詳象賞醤鉦鍾鐘障鞘上丈丞乗冗剰城場壌嬢常情擾条杖浄状畳穣蒸譲醸錠嘱埴飾
拭植殖燭織職色触食蝕辱尻伸信侵唇娠寝審心慎振新晋森榛浸深申疹真神秦紳臣芯薪親診身辛進針震人仁刃塵壬尋甚尽腎訊迅陣靭笥諏須酢図厨逗吹垂帥推水炊睡粋翠衰遂酔錐錘随瑞髄崇嵩数枢趨雛据杉椙菅頗雀裾
澄摺寸世瀬畝是凄制勢姓征性成政整星晴棲栖正清牲生盛精聖声製西誠誓請逝醒青静斉税脆隻席惜戚斥昔析石積籍績脊責赤跡蹟碩切拙接摂折設窃節説雪絶舌蝉仙先千占宣専尖川戦扇撰栓栴泉浅洗染潜煎煽旋穿箭線
繊羨腺舛船薦詮賎践選遷銭銑閃鮮前善漸然全禅繕膳糎噌塑岨措曾曽楚狙疏疎礎祖租粗素組蘇訴阻遡鼠僧創双叢倉喪壮奏爽宋層匝惣想捜掃挿掻操早曹巣槍槽漕燥争痩相窓糟総綜聡草荘葬蒼藻装走送遭鎗霜騒像増憎
臓蔵贈造促側則即息捉束測足速俗属賊族続卒袖其揃存孫尊損村遜他多太汰詑唾堕妥惰打柁舵楕陀駄騨体堆対耐岱帯待怠態戴替泰滞胎腿苔袋貸退逮隊黛鯛代台大第醍題鷹滝瀧卓啄宅托択拓沢濯琢託鐸濁諾茸凧蛸只
叩但達辰奪脱巽竪辿棚谷狸鱈樽誰丹単嘆坦担探旦歎淡湛炭短端箪綻耽胆蛋誕鍛団壇弾断暖檀段男談値知地弛恥智池痴稚置致蜘遅馳築畜竹筑蓄逐秩窒茶嫡着中仲宙忠抽昼柱注虫衷註酎鋳駐樗瀦猪苧著貯丁兆凋喋寵
帖帳庁弔張彫徴懲挑暢朝潮牒町眺聴脹腸蝶調諜超跳銚長頂鳥勅捗直朕沈珍賃鎮陳津墜椎槌追鎚痛通塚栂掴槻佃漬柘辻蔦綴鍔椿潰坪壷嬬紬爪吊釣鶴亭低停偵剃貞呈堤定帝底庭廷弟悌抵挺提梯汀碇禎程締艇訂諦蹄逓
邸鄭釘鼎泥摘擢敵滴的笛適鏑溺哲徹撤轍迭鉄典填天展店添纏甜貼転顛点伝殿澱田電兎吐堵塗妬屠徒斗杜渡登菟賭途都鍍砥砺努度土奴怒倒党冬凍刀唐塔塘套宕島嶋悼投搭東桃梼棟盗淘湯涛灯燈当痘祷等答筒糖統到
董蕩藤討謄豆踏逃透鐙陶頭騰闘働動同堂導憧撞洞瞳童胴萄道銅峠鴇匿得徳涜特督禿篤毒独読栃橡凸突椴届鳶苫寅酉瀞噸屯惇敦沌豚遁頓呑曇鈍奈那内乍凪薙謎灘捺鍋楢馴縄畷南楠軟難汝二尼弐迩匂賑肉虹廿日乳入
如尿韮任妊忍認濡禰祢寧葱猫熱年念捻撚燃粘乃廼之埜嚢悩濃納能脳膿農覗蚤巴把播覇杷波派琶破婆罵芭馬俳廃拝排敗杯盃牌背肺輩配倍培媒梅楳煤狽買売賠陪這蝿秤矧萩伯剥博拍柏泊白箔粕舶薄迫曝漠爆縛莫駁麦
函箱硲箸肇筈櫨幡肌畑畠八鉢溌発醗髪伐罰抜筏閥鳩噺塙蛤隼伴判半反叛帆搬斑板氾汎版犯班畔繁般藩販範釆煩頒飯挽晩番盤磐蕃蛮匪卑否妃庇彼悲扉批披斐比泌疲皮碑秘緋罷肥被誹費避非飛樋簸備尾微枇毘琵眉美
鼻柊稗匹疋髭彦膝菱肘弼必畢筆逼桧姫媛紐百謬俵彪標氷漂瓢票表評豹廟描病秒苗錨鋲蒜蛭鰭品彬斌浜瀕貧賓頻敏瓶不付埠夫婦富冨布府怖扶敷斧普浮父符腐膚芙譜負賦赴阜附侮撫武舞葡蕪部封楓風葺蕗伏副復幅服
福腹複覆淵弗払沸仏物鮒分吻噴墳憤扮焚奮粉糞紛雰文聞丙併兵塀幣平弊柄並蔽閉陛米頁僻壁癖碧別瞥蔑箆偏変片篇編辺返遍便勉娩弁鞭保舗鋪圃捕歩甫補輔穂募墓慕戊暮母簿菩倣俸包呆報奉宝峰峯崩庖抱捧放方朋
法泡烹砲縫胞芳萌蓬蜂褒訪豊邦鋒飽鳳鵬乏亡傍剖坊妨帽忘忙房暴望某棒冒紡肪膨謀貌貿鉾防吠頬北僕卜墨撲朴牧睦穆釦勃没殆堀幌奔本翻凡盆摩磨魔麻埋妹昧枚毎哩槙幕膜枕鮪柾鱒桝亦俣又抹末沫迄侭繭麿万慢満
漫蔓味未魅巳箕岬密蜜湊蓑稔脈妙粍民眠務夢無牟矛霧鵡椋婿娘冥名命明盟迷銘鳴姪牝滅免棉綿緬面麺摸模茂妄孟毛猛盲網耗蒙儲木黙目杢勿餅尤戻籾貰問悶紋門匁也冶夜爺耶野弥矢厄役約薬訳躍靖柳薮鑓愉愈油癒
諭輸唯佑優勇友宥幽悠憂揖有柚湧涌猶猷由祐裕誘遊邑郵雄融夕予余与誉輿預傭幼妖容庸揚揺擁曜楊様洋溶熔用窯羊耀葉蓉要謡踊遥陽養慾抑欲沃浴翌翼淀羅螺裸来莱頼雷洛絡落酪乱卵嵐欄濫藍蘭覧利吏履李梨理璃
痢裏裡里離陸律率立葎掠略劉流溜琉留硫粒隆竜龍侶慮旅虜了亮僚両凌寮料梁涼猟療瞭稜糧良諒遼量陵領力緑倫厘林淋燐琳臨輪隣鱗麟瑠塁涙累類令伶例冷励嶺怜玲礼苓鈴隷零霊麗齢暦歴列劣烈裂廉恋憐漣煉簾練聯
蓮連錬呂魯櫓炉賂路露労婁廊弄朗楼榔浪漏牢狼篭老聾蝋郎六麓禄肋録論倭和話歪賄脇惑枠鷲亙亘鰐詫藁蕨椀湾碗腕
`

	const translatedTable = `
Afraid of foreign migration, dimensional weft, stomach withering, said to be illegal medicine, Jinghaiyuyuyujiyiyiyiyiyiyicitaoqiaoyinyanyan members are induced to drink obscene Yinyin due to marriage
Yin Yin, rhyme, right Yu, black feathers, roundabout rain, pelican peeking at the ugly mortar, vortex, peeping, Wei eel, stable, melon, leaping clouds, clouds flowing, bait, Rui palace, baby shadow, shadow, dragging, Yongying, Yingying, Yingyingwei. Chanting the liquid epidemic, benefiting the station, Yueyue, Yueyue, Yanyuan, Yanyuan, banquet, delay, support, Yan Yanyan, Yanyan, ape, 牶苑薗yuan, lead mandarin duck, in the dirty nephew, Yangao, go to Xian
Wang Wang Heng Ou beat Wang Weng's coat 鴴鴴Huanggang Chong Diyiwu Yiyiju barrel oyi I'm unloading grace and gentle sound. Crop flowers, eggplant, lotus fruit, shrimp class, Huohuo, Gaoxia, mosquito, O'er, my tooth painting, lying bud moth, Heya, hungry driving introduction, will be solved, and the monster will be regretted and changed.
The armored carcass, frogs, persimmons, oysters, hooks, and shells were all found in the sea. Compared with Guo Ge, learn from Yue's forehead and jaw, hang Li, 樫橿梶鳛, cut and drink, just live, slide, kudzu, brown, bonito leaves, birch, tree, pocket, pu cauldron, sickle, duck, cypress, and Xuan
Porridge, cutting, tiles, drying, talking about crowns and colds, looking at the wind and waves, summoning the officials, punishing the officials, feeling used to it, regretting, changing the courage, tangerine, coffin, sweat, Hanjian, Huangan, supervision, pole, simple, slow, fouhan, livership, Wankan, and Guanhui. Jianjian Xianguan, Hanguanguanwanwanwanwan, playing with cancer eyes, rocks, false goose, stubborn faces, wishing to Qi, dangerous, happy, Ji Qi, Ji Qi, Xi, Ji Ji, waving the flag, discarding the game at the scheduled time.
Machine, Yi, Qi, Qi, Qi, Qi, Ji, Ji, Ji, Ji, Ji, Ji, Hui, Hungry, Riding, Ghost, False ritual, Prostitute, Suitable for play, trick, deception, doubt, righteousness, ant friendship, chrysanthemum, juji, eat, orange, anvil, pestle, millet, but guest The foot abuses Niqiu for a long time, and the uterine bow is sucked for first aid. The dead beg to weep and moxibustion ball to find out the poor level and give it to the old cow to live in the giant.
Kuangqing called Qiao Jingxia to be strong, timid and respectful, to teach the bridge situation, to be mad and narrow, to straighten the breasts and ribs, to raise the mirror, to be startled, to congeal, to congeal, to be honest, to be honest. Xin Qin Qin forbidden birds tendons tight celery fungus Jin Jin close Jin Yin Yin Jiu Jiu District dog Jiu moment bitter body 駆駈horse tool foolish Yu Gong empty occasional residence encounter corner string comb 鏶 crumbs Qu
Digging a hole, a shoe, a bridle, a bear, a chestnut, a chestnut, a shovel, a shovel, a wormwood, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a shovel, a hoe, a hoe Yong Ji Yi Jing Yi Jing Yun Ying Whale Drama Halberd Ji Jie Qiu Jue Jie Jie Xue Jie Yue Jian Jian Jian Jian Xuan Huo Jian Su Jian Xian Hang Fist Scroll
検権珩狗 Presents the grinding inkstone and silk Hakama Hu, wild tiger, praise, cross cobalt, hire Gu Guwu, Wu Wu, Wu Tuo, the queen, the emperor realizes Wu, Hu, Qiyu, protects the wine by mistake, begs for the carp, pays the best, Hou Hou, Xingguang, public servant, hooks thick mouth.
Back throat pit dirt good hole Xiao Hong Gong Qiao Lane Xing Geng Kang Hong Heng panic resist arrest attack Ang Huang Geng Hang school stem structure Jiang Hong Hao Gang ditch Jia Huang hard draft chaff Hong Hong twist gang plowing test Ken humerus plaster sailing waste line balance lecture tribute purchase suburb Fermentation, steel, steel, Xiang Xiang, Gao Honggang, robbery, trench, torture, Hao Hao, and Qu Ke carved to sue the country, cool swan, black prison, wet waist steamer, sudden trance, bone 狛込
At this moment, I am trapped in my marriage, I am in a coma, my roots are mixed with traces, my soul is cyanotic, my soul is confused, I am in trouble, I am in trouble, I am lying on my back, I am lying on my back, I am lying on my back, I am sitting on my back, I am debt-ridden, I am in trouble, my wife, Zai Caicai, is in trouble. Harvesting the rhinoceros and sacrificial vegetables, cutting the fine vegetables, cutting the material, the material, the wealth, the sakaki, the sakai, the sakaki, the heron, the heron, the sakaki, the sage, the heron, the sakaki, the heron, the sakura, the sakura, the sakura, the sakura, the salmon, the spoon, the brush
Check, wipe, and kill Sa Yong Gao, 8, 4, 7, and 10 sharks. Three umbrellas, ginseng and mountain. Scattered handbags, Canshan property calculations, silkworms, praising sour meals, beheading the temporarily disabled officials, waiting for the governor, Shi Si, the fourth official, the first sister Zizi Corpse City Master Zhi Si Zhi Zhi Si Shi Shi Zhi Zhi Zhi Shi Shi Zhi Private Silk Paper Purple Limb Fat Zhi Shi Ci Poetry Examination Zhi Zi Zi Grant Female Feeding
Second Zi Zhi Er Xi Hemorrhoids Magnetism Shows and Er Er Shi Ci Xi Deer Style Knows Zhu Zhu Axis Shishi Shizuku Qichi Zhi Loss Jealousy Room Knows Damp Lacquer Disease Quality 実蔀 Xiao Si Chai Zhi Repeated Rui Rui Jie Shao Writes Sheshe Pardon and Xie Boils Shasha The person thanks the car to hide the evil spirits of the snake. Borrow the spoon, ruler, ladle, burn the noble wine, and drink the pewter. If he is lonely, he offends the master. He takes the guard Zhu Shu, hunts for pearls, grows swollen wine, and the Confucian scholar receives the longevity gift from the tree. He needs to be imprisoned to collect the Zhou Dynasty.
Zong Jiuzhou repairs sorrow, picks up the continents, shows off the autumn, learns the smelly boat, collects the people, attacks the enemy, compiles the Zhou chieftain's reward, gathers the ugly and even lives, fills the ten Rong soft juice, the heavy gun, the uncle, the long stay, Shu Zhu, the shrunken school, is familiar with the art, tells the Jun Jun Chun Shunjun Junxun Xunzhen martyred Chunjun Rundun pure patrol followed Chunshun Chushuo Shushu Zhuxu Xu Department Shushushu Zhushushuxu women's sequence Xu Shu hoeing and compensation for injuries
Shengjiang was promoted to summon the sentry merchant to sing and taste concubines and prostitutes at night. Chang Jiang Jiao Chong Shang Litigation Edict Detailed Image Appreciation of the Clock, Bell, Bell, Barrier and Sheath, Zhang Cheng Riding in the City Market, Lady and Maidens Always Frequently Annoying, Bar Staff, Pure Shape, Steamed Powder, Ingots, Instructions and Ornaments
Wipe the plant, weave, color, touch food, corrode, humiliate, extend the letter, invade the lips, pregnancy, sleep, carefully vibrate the new forest hazel, soak deeply, apply rash, the true god Qin Shenchen, personally diagnose the body, puncture the needle, shock the human heart, blade dust, Ren Xunshi When the kidneys are exhausted, the news is fast. The toughness is strong.
The world is clear, the mother is sad, the power is surnamed, the government is established, the star Qingqiqi is pure, the life is prosperous, the holy voice is made, the west is sincere, swear to die, wake up, Qingjing, tax, just sit, cherish, rebuke the past, analyze the stone, accumulate achievements The spine is full of scars, the traces are huge, the clumsiness is connected, the stealing section is said, the snow is gone, the tongue is cicada, the immortal is Xian Qianzhan, Xuan Shang Jian, Sichuan Fan, Shuan Zhuanquan, shallow washed, dyed, latent fried, stirred, swirled and pierced the arrow line.
繊曽 gland 舛舛舛曰困拷 practices selection and relocation 銭 mill flash fresh qian shan gradually quan chan repair meal 糎獌sculpture 娨cuo 曽潽潽潽曽氛笎 basin ancestral rent rough element group su suing to stop the squirrel monk creating a double cluster warehouse The funeral was solemn, the Song Dynasty was full of joy, and the Song Dynasty wanted to sweep and insert it, Cao Cao Cao Si's gun slot was dry, and he was fighting for the chancellor.
The gift of 蓓蔵, the promotion of the side, is about to stop, catch the bundle, measure the speed, and belong to the thieves clan. Waiting for the idle state, wearing Taiyang, fetal leg moss bag, borrowing and retreating, catching the team, Dai Hai, on behalf of National Taiwan University.
Knock but Dachen seizes the Xun vertical shed valley raccoon cod bottle who Dan sighs Tandan Tandan sighs light Zhan charcoal short end Tantan blooms Dan Dan Dan forged forged altar 団 Tan Duan Nan talks about the value of knowing the ground and relaxing the shame Zhichi Stupid and childlike setting causes spiders to run wildly and build livestock and bamboos to build storage.
The tents are hung, the sculptures are hung, the eagles are raised, the chaochuang is watched, the tide is turned out, the butterfly is transferred, the spies are jumped, the long-headed bird is edicted, I am buried in Chenjin, the vertebrae are dropped, the hammer is chased, the hammer is painful, the tomb is stained昘軻茑贷洷洿洷平壷嬬第paw 笴第嬬洬狠狠 investigation
Di Zheng nails the cauldron mud picks up the enemy's drop of the flute suitable dysprosium drowning Zhe Che withdraws the track the iron code fills the sky the exhibition shop adds the sweet stickers the dimpling point the temple Yododa the electric chot vomits the jealous butcher fights the Dudu climbs the gambling The road is polished, Nudu, the slaves are angry, the party is frozen, the knife is frozen in Tangtangtang, Taodangdao Island is mourned, and the Dongtaoying Building is stolen, the soup is washed, the lantern is lit, and the pimples are prayed, waiting for the sugar cane to arrive.
Dong Dangteng begged for a copy of the bean and escaped through the stirrup. You Chengton Tun Dun Dun Dun Chao Dolphin Tun Dun Dun Duo Nai Na Nei Zha Nag Nai Nai Na Na Na Guo Nara Tamao Wan Nan Nan Soft Hard Ru Erni Ni Nai Nai Relieving Meat Hong Twenty Days Breast In
Such as urinating, letting pregnant women bear to recognize wetness, you, Ning, onion, cat, hot year, thinking, twisting, burning and sticky, but it is the wild thing, it is rich, it can ward off pus, it can ward off pus, it can ward off fleas, it can cause irritation Cup cup brand back lung generation with double pei matchmaker Mei Mei coal 狋buy 売compensation to accompany this clover scale 秧萐博 peel off bopai baibo white foil meal ship thin forced exposure desert explosion tied Mo pu wheat
Letter box, chopsticks, chopsticks, flags, flags, and flags In the evening, the barbarian bandits were humble, but the concubines protected their sorrows.
Nose, weeds, mustaches, knees, rhombs, elbows, bibibi, pens, cypress, cypress, new, 100-year-olds, Biaobiao, ice-floating scoops, table comments, leopard temple, description of disease, seconds seedlings, anchor, garlic, leech fins, product binbin, verge of poverty, bin frequency Pingbufubu couple Fuhanbu Fufu Fu Fu Fu Pu Fu Fu Fu Fufu Pu Fu Fu went to Fufu Fu Fu Wu Wu Wu Pu Wu Bu Feng Feng Feng Pei Lotus Fu Fu Fu Fu Fu Fufu
The belly is covered with abyss, boiling things are boiled, the things are divided, the grave is kissed, the grave is filled with anger, the powder is mixed with excrement, the atmosphere is full of literature and knowledge, the soldiers are defeated, the currency is flat, the handle is closed, the rice page is hidden in a remote wall, and the seal is scorned. After editing and re-reading the chapter, I reluctantly gave birth, Benbian Baopu Pupu caught the steps, repaired the auxiliary work, raised the tomb, Mu Wumu's mother book, Bodhisattva, and paid my salary to stay.
French soaked, cooked, seamed, fragrant, fragrant, bee praised, visited Feng Bangfeng, Fengpeng was exhausted, died next to the cross section, blocked the hat, forgot about the busy room, looked at a stick, spun fat, swollen, Maobo, barked, Beipu, Mopu, Pu Mumu. Mu Kobo is dead, and he is running around, rubbing the magic hemp in the mortal pot, burying the girl, every mile, the curtain film pillow, the tuna, the trout, the 桝, and the moth, the foam, the cocoon, the cocoon, the cocoon, the cocoon, the cocoon.
The smell of vines is not charming. The secret honey of the Jiji Cape is the coir and the pulse of the wonderful rice. The people sleep in dreams without Mou spears. Mist parrots. Starling. Mao Meng blind net consumes Meng Chu Mu 黙目杢不吃饼 Youjuan Fen贳 asks dull pattern Menjiaoyeyeye Yeye Yiya's service about 薬訳yue Jingliu 薮鑓please get better and you get better
The message is lost, only the excellent and brave friends are forgiven and you are worried, there is a pomelo gushing out, and you are still smelting. Yang Yao Ye Rong wants to go to the far yang to nourish the desire and suppress the desire. Fertile bath. Yiyi Lake.
Li Lili Li Lu Lv led Li Qi to plunder Liu Liuliu Liuliu Sulfur Grain Long Longlong Lvlu captured Liang Liao Liang Liang Liang Liang Liang Liang Liang Liang Liang Liao Liang Ling Lvlun Li Linlin Lin Lin Ling Lin Lin Lin 涙 tired class Ling Ling example Leng Li Ling Ling Li Ling Ling Li Ling Li Ling Li Ling Li Ling Li You Lie Lian Lian Lian Lian Lianlian Lianlian
Lotus, Lulu, Lu, Lu, Lu, Lu, Lu, Lu, Lou, corridor, Nonglang, Lang, Lang, leak, prison, wolf, old deaf man, Liulu, Lulilu, Discussing Japanese peace, distortion, bribery, threats, vultures, crocodiles, ferns, bowls, and wrists
`
	for _, c := range translatedTable {
		if c == '\n' {
			continue
		}
		jaKanjis = append(jaKanjis, c)
	}
}

var (
	mplusFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	counter        int
	kanjiText      string
	kanjiTextColor color.RGBA
}

func (g *Game) Update() error {
	// Change the text color for each second.
	if g.counter%ebiten.TPS() == 0 {
		g.kanjiText = ""
		for j := 0; j < 6; j++ {
			for i := 0; i < 12; i++ {
				g.kanjiText += string(jaKanjis[rand.Intn(len(jaKanjis))])
			}
			g.kanjiText += "\n"
		}

		g.kanjiTextColor.R = 0x80 + uint8(rand.Intn(0x7f))
		g.kanjiTextColor.G = 0x80 + uint8(rand.Intn(0x7f))
		g.kanjiTextColor.B = 0x80 + uint8(rand.Intn(0x7f))
		g.kanjiTextColor.A = 0xff
	}
	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		normalFontSize = 24
		bigFontSize    = 48
	)

	const x = 20

	// Draw info
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, 20)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, op)

	// Draw the sample text
	op = &text.DrawOptions{}
	op.GeoM.Translate(x, 60)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, sampleText, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   normalFontSize,
	}, op)

	// Draw Kanji text lines
	op = &text.DrawOptions{}
	op.GeoM.Translate(x, 110)
	op.ColorScale.ScaleWithColor(g.kanjiTextColor)
	op.LineSpacing = bigFontSize * 1.2
	text.Draw(screen, g.kanjiText, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   bigFontSize,
	}, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Font (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
