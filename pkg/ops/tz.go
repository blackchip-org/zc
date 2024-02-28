package ops

var TimeZones = map[string]string {
	"tz.abidjan.africa": "[Africa/Abidjan]",
	"tz.accra.africa": "[Africa/Accra]",
	"tz.acre.brazil": "[Brazil/Acre]",
	"tz.act.australia": "[Australia/ACT]",
	"tz.adak.america": "[America/Adak]",
	"tz.addis-ababa.africa": "[Africa/Addis_Ababa]",
	"tz.adelaide.australia": "[Australia/Adelaide]",
	"tz.aden.asia": "[Asia/Aden]",
	"tz.alaska.us": "[US/Alaska]",
	"tz.aleutian.us": "[US/Aleutian]",
	"tz.algiers.africa": "[Africa/Algiers]",
	"tz.almaty.asia": "[Asia/Almaty]",
	"tz.amman.asia": "[Asia/Amman]",
	"tz.amsterdam.europe": "[Europe/Amsterdam]",
	"tz.anadyr.asia": "[Asia/Anadyr]",
	"tz.anchorage.america": "[America/Anchorage]",
	"tz.andorra.europe": "[Europe/Andorra]",
	"tz.anguilla.america": "[America/Anguilla]",
	"tz.antananarivo.indian": "[Indian/Antananarivo]",
	"tz.antigua.america": "[America/Antigua]",
	"tz.apia.pacific": "[Pacific/Apia]",
	"tz.aqtau.asia": "[Asia/Aqtau]",
	"tz.aqtobe.asia": "[Asia/Aqtobe]",
	"tz.araguaina.america": "[America/Araguaina]",
	"tz.arizona.us": "[US/Arizona]",
	"tz.aruba.america": "[America/Aruba]",
	"tz.ashgabat.asia": "[Asia/Ashgabat]",
	"tz.ashkhabad.asia": "[Asia/Ashkhabad]",
	"tz.asmara.africa": "[Africa/Asmara]",
	"tz.asmera.africa": "[Africa/Asmera]",
	"tz.astrakhan.europe": "[Europe/Astrakhan]",
	"tz.asuncion.america": "[America/Asuncion]",
	"tz.athens.europe": "[Europe/Athens]",
	"tz.atikokan.america": "[America/Atikokan]",
	"tz.atka.america": "[America/Atka]",
	"tz.atlantic.canada": "[Canada/Atlantic]",
	"tz.atyrau.asia": "[Asia/Atyrau]",
	"tz.auckland.pacific": "[Pacific/Auckland]",
	"tz.australia.north": "[Australia/North]",
	"tz.australia.west": "[Australia/West]",
	"tz.azores.atlantic": "[Atlantic/Azores]",
	"tz.baghdad.asia": "[Asia/Baghdad]",
	"tz.bahia-banderas.america": "[America/Bahia_Banderas]",
	"tz.bahia.america": "[America/Bahia]",
	"tz.bahrain.asia": "[Asia/Bahrain]",
	"tz.bajanorte.mexico": "[Mexico/BajaNorte]",
	"tz.bajasur.mexico": "[Mexico/BajaSur]",
	"tz.baku.asia": "[Asia/Baku]",
	"tz.bamako.africa": "[Africa/Bamako]",
	"tz.bangkok.asia": "[Asia/Bangkok]",
	"tz.bangui.africa": "[Africa/Bangui]",
	"tz.banjul.africa": "[Africa/Banjul]",
	"tz.barbados.america": "[America/Barbados]",
	"tz.barnaul.asia": "[Asia/Barnaul]",
	"tz.beirut.asia": "[Asia/Beirut]",
	"tz.belem.america": "[America/Belem]",
	"tz.belfast.europe": "[Europe/Belfast]",
	"tz.belgrade.europe": "[Europe/Belgrade]",
	"tz.belize.america": "[America/Belize]",
	"tz.berlin.europe": "[Europe/Berlin]",
	"tz.bermuda.atlantic": "[Atlantic/Bermuda]",
	"tz.beulah.north-dakota.america": "[America/North_Dakota/Beulah]",
	"tz.bishkek.asia": "[Asia/Bishkek]",
	"tz.bissau.africa": "[Africa/Bissau]",
	"tz.blanc-sablon.america": "[America/Blanc-Sablon]",
	"tz.blantyre.africa": "[Africa/Blantyre]",
	"tz.boa-vista.america": "[America/Boa_Vista]",
	"tz.bogota.america": "[America/Bogota]",
	"tz.boise.america": "[America/Boise]",
	"tz.bougainville.pacific": "[Pacific/Bougainville]",
	"tz.bratislava.europe": "[Europe/Bratislava]",
	"tz.brazil.east": "[Brazil/East]",
	"tz.brazil.west": "[Brazil/West]",
	"tz.brazzaville.africa": "[Africa/Brazzaville]",
	"tz.brisbane.australia": "[Australia/Brisbane]",
	"tz.broken-hill.australia": "[Australia/Broken_Hill]",
	"tz.brunei.asia": "[Asia/Brunei]",
	"tz.brussels.europe": "[Europe/Brussels]",
	"tz.bucharest.europe": "[Europe/Bucharest]",
	"tz.budapest.europe": "[Europe/Budapest]",
	"tz.buenos-aires.america": "[America/Buenos_Aires]",
	"tz.buenos-aires.argentina.america": "[America/Argentina/Buenos_Aires]",
	"tz.bujumbura.africa": "[Africa/Bujumbura]",
	"tz.busingen.europe": "[Europe/Busingen]",
	"tz.cairo.africa": "[Africa/Cairo]",
	"tz.calcutta.asia": "[Asia/Calcutta]",
	"tz.cambridge-bay.america": "[America/Cambridge_Bay]",
	"tz.campo-grande.america": "[America/Campo_Grande]",
	"tz.canada.atlantic": "[Canada/Atlantic]",
	"tz.canada.eastern": "[Canada/Eastern]",
	"tz.canada.pacific": "[Canada/Pacific]",
	"tz.canary.atlantic": "[Atlantic/Canary]",
	"tz.canberra.australia": "[Australia/Canberra]",
	"tz.cancun.america": "[America/Cancun]",
	"tz.cape-verde.atlantic": "[Atlantic/Cape_Verde]",
	"tz.caracas.america": "[America/Caracas]",
	"tz.casablanca.africa": "[Africa/Casablanca]",
	"tz.casey.antarctica": "[Antarctica/Casey]",
	"tz.catamarca.america": "[America/Catamarca]",
	"tz.catamarca.argentina.america": "[America/Argentina/Catamarca]",
	"tz.cayenne.america": "[America/Cayenne]",
	"tz.cayman.america": "[America/Cayman]",
	"tz.center.north-dakota.america": "[America/North_Dakota/Center]",
	"tz.central.canada": "[Canada/Central]",
	"tz.central.us": "[US/Central]",
	"tz.cet": "[CET]",
	"tz.ceuta.africa": "[Africa/Ceuta]",
	"tz.chagos.indian": "[Indian/Chagos]",
	"tz.chatham.pacific": "[Pacific/Chatham]",
	"tz.chicago.america": "[America/Chicago]",
	"tz.chihuahua.america": "[America/Chihuahua]",
	"tz.chisinau.europe": "[Europe/Chisinau]",
	"tz.chita.asia": "[Asia/Chita]",
	"tz.choibalsan.asia": "[Asia/Choibalsan]",
	"tz.chongqing.asia": "[Asia/Chongqing]",
	"tz.christmas.indian": "[Indian/Christmas]",
	"tz.chungking.asia": "[Asia/Chungking]",
	"tz.chuuk.pacific": "[Pacific/Chuuk]",
	"tz.ciudad-juarez.america": "[America/Ciudad_Juarez]",
	"tz.cocos.indian": "[Indian/Cocos]",
	"tz.colombo.asia": "[Asia/Colombo]",
	"tz.comodrivadavia.argentina.america": "[America/Argentina/ComodRivadavia]",
	"tz.comoro.indian": "[Indian/Comoro]",
	"tz.conakry.africa": "[Africa/Conakry]",
	"tz.continental.chile": "[Chile/Continental]",
	"tz.copenhagen.europe": "[Europe/Copenhagen]",
	"tz.coral-harbour.america": "[America/Coral_Harbour]",
	"tz.cordoba.america": "[America/Cordoba]",
	"tz.cordoba.argentina.america": "[America/Argentina/Cordoba]",
	"tz.costa-rica.america": "[America/Costa_Rica]",
	"tz.creston.america": "[America/Creston]",
	"tz.cst6cdt": "[CST6CDT]",
	"tz.cuba": "[Cuba]",
	"tz.cuiaba.america": "[America/Cuiaba]",
	"tz.curacao.america": "[America/Curacao]",
	"tz.currie.australia": "[Australia/Currie]",
	"tz.dacca.asia": "[Asia/Dacca]",
	"tz.dakar.africa": "[Africa/Dakar]",
	"tz.damascus.asia": "[Asia/Damascus]",
	"tz.danmarkshavn.america": "[America/Danmarkshavn]",
	"tz.dar-es-salaam.africa": "[Africa/Dar_es_Salaam]",
	"tz.darwin.australia": "[Australia/Darwin]",
	"tz.davis.antarctica": "[Antarctica/Davis]",
	"tz.dawson-creek.america": "[America/Dawson_Creek]",
	"tz.dawson.america": "[America/Dawson]",
	"tz.denoronha.brazil": "[Brazil/DeNoronha]",
	"tz.denver.america": "[America/Denver]",
	"tz.detroit.america": "[America/Detroit]",
	"tz.dhaka.asia": "[Asia/Dhaka]",
	"tz.dili.asia": "[Asia/Dili]",
	"tz.djibouti.africa": "[Africa/Djibouti]",
	"tz.dominica.america": "[America/Dominica]",
	"tz.douala.africa": "[Africa/Douala]",
	"tz.dubai.asia": "[Asia/Dubai]",
	"tz.dublin.europe": "[Europe/Dublin]",
	"tz.dumontdurville.antarctica": "[Antarctica/DumontDUrville]",
	"tz.dushanbe.asia": "[Asia/Dushanbe]",
	"tz.east-indiana.us": "[US/East-Indiana]",
	"tz.east.brazil": "[Brazil/East]",
	"tz.easter.pacific": "[Pacific/Easter]",
	"tz.easterisland.chile": "[Chile/EasterIsland]",
	"tz.eastern.canada": "[Canada/Eastern]",
	"tz.eastern.us": "[US/Eastern]",
	"tz.edmonton.america": "[America/Edmonton]",
	"tz.eet": "[EET]",
	"tz.efate.pacific": "[Pacific/Efate]",
	"tz.egypt": "[Egypt]",
	"tz.eire": "[Eire]",
	"tz.eirunepe.america": "[America/Eirunepe]",
	"tz.el-aaiun.africa": "[Africa/El_Aaiun]",
	"tz.el-salvador.america": "[America/El_Salvador]",
	"tz.enderbury.pacific": "[Pacific/Enderbury]",
	"tz.ensenada.america": "[America/Ensenada]",
	"tz.est": "[EST]",
	"tz.est5edt": "[EST5EDT]",
	"tz.eucla.australia": "[Australia/Eucla]",
	"tz.faeroe.atlantic": "[Atlantic/Faeroe]",
	"tz.fakaofo.pacific": "[Pacific/Fakaofo]",
	"tz.famagusta.asia": "[Asia/Famagusta]",
	"tz.faroe.atlantic": "[Atlantic/Faroe]",
	"tz.fiji.pacific": "[Pacific/Fiji]",
	"tz.fort-nelson.america": "[America/Fort_Nelson]",
	"tz.fort-wayne.america": "[America/Fort_Wayne]",
	"tz.fortaleza.america": "[America/Fortaleza]",
	"tz.freetown.africa": "[Africa/Freetown]",
	"tz.funafuti.pacific": "[Pacific/Funafuti]",
	"tz.gaborone.africa": "[Africa/Gaborone]",
	"tz.galapagos.pacific": "[Pacific/Galapagos]",
	"tz.gambier.pacific": "[Pacific/Gambier]",
	"tz.gaza.asia": "[Asia/Gaza]",
	"tz.gb": "[GB]",
	"tz.gb-eire": "[GB-Eire]",
	"tz.general.mexico": "[Mexico/General]",
	"tz.gibraltar.europe": "[Europe/Gibraltar]",
	"tz.glace-bay.america": "[America/Glace_Bay]",
	"tz.gmt": "[GMT]",
	"tz.gmt+0": "[GMT+0]",
	"tz.gmt+1": "[Etc/GMT+1]",
	"tz.gmt+10": "[Etc/GMT+10]",
	"tz.gmt+11": "[Etc/GMT+11]",
	"tz.gmt+12": "[Etc/GMT+12]",
	"tz.gmt+2": "[Etc/GMT+2]",
	"tz.gmt+3": "[Etc/GMT+3]",
	"tz.gmt+4": "[Etc/GMT+4]",
	"tz.gmt+5": "[Etc/GMT+5]",
	"tz.gmt+6": "[Etc/GMT+6]",
	"tz.gmt+7": "[Etc/GMT+7]",
	"tz.gmt+8": "[Etc/GMT+8]",
	"tz.gmt+9": "[Etc/GMT+9]",
	"tz.gmt-0": "[GMT-0]",
	"tz.gmt-1": "[Etc/GMT-1]",
	"tz.gmt-10": "[Etc/GMT-10]",
	"tz.gmt-11": "[Etc/GMT-11]",
	"tz.gmt-12": "[Etc/GMT-12]",
	"tz.gmt-13": "[Etc/GMT-13]",
	"tz.gmt-14": "[Etc/GMT-14]",
	"tz.gmt-2": "[Etc/GMT-2]",
	"tz.gmt-3": "[Etc/GMT-3]",
	"tz.gmt-4": "[Etc/GMT-4]",
	"tz.gmt-5": "[Etc/GMT-5]",
	"tz.gmt-6": "[Etc/GMT-6]",
	"tz.gmt-7": "[Etc/GMT-7]",
	"tz.gmt-8": "[Etc/GMT-8]",
	"tz.gmt-9": "[Etc/GMT-9]",
	"tz.gmt0": "[GMT0]",
	"tz.godthab.america": "[America/Godthab]",
	"tz.goose-bay.america": "[America/Goose_Bay]",
	"tz.grand-turk.america": "[America/Grand_Turk]",
	"tz.greenwich": "[Greenwich]",
	"tz.grenada.america": "[America/Grenada]",
	"tz.guadalcanal.pacific": "[Pacific/Guadalcanal]",
	"tz.guadeloupe.america": "[America/Guadeloupe]",
	"tz.guam.pacific": "[Pacific/Guam]",
	"tz.guatemala.america": "[America/Guatemala]",
	"tz.guayaquil.america": "[America/Guayaquil]",
	"tz.guernsey.europe": "[Europe/Guernsey]",
	"tz.guyana.america": "[America/Guyana]",
	"tz.halifax.america": "[America/Halifax]",
	"tz.harare.africa": "[Africa/Harare]",
	"tz.harbin.asia": "[Asia/Harbin]",
	"tz.havana.america": "[America/Havana]",
	"tz.hawaii.us": "[US/Hawaii]",
	"tz.hebron.asia": "[Asia/Hebron]",
	"tz.helsinki.europe": "[Europe/Helsinki]",
	"tz.hermosillo.america": "[America/Hermosillo]",
	"tz.ho-chi-minh.asia": "[Asia/Ho_Chi_Minh]",
	"tz.hobart.australia": "[Australia/Hobart]",
	"tz.hong-kong.asia": "[Asia/Hong_Kong]",
	"tz.hongkong": "[Hongkong]",
	"tz.honolulu.pacific": "[Pacific/Honolulu]",
	"tz.hovd.asia": "[Asia/Hovd]",
	"tz.hst": "[HST]",
	"tz.iceland": "[Iceland]",
	"tz.indiana-starke.us": "[US/Indiana-Starke]",
	"tz.indianapolis.america": "[America/Indianapolis]",
	"tz.indianapolis.indiana.america": "[America/Indiana/Indianapolis]",
	"tz.inuvik.america": "[America/Inuvik]",
	"tz.iqaluit.america": "[America/Iqaluit]",
	"tz.iran": "[Iran]",
	"tz.irkutsk.asia": "[Asia/Irkutsk]",
	"tz.isle-of-man.europe": "[Europe/Isle_of_Man]",
	"tz.israel": "[Israel]",
	"tz.istanbul.asia": "[Asia/Istanbul]",
	"tz.istanbul.europe": "[Europe/Istanbul]",
	"tz.jakarta.asia": "[Asia/Jakarta]",
	"tz.jamaica": "[Jamaica]",
	"tz.jamaica.america": "[America/Jamaica]",
	"tz.jan-mayen.atlantic": "[Atlantic/Jan_Mayen]",
	"tz.japan": "[Japan]",
	"tz.jayapura.asia": "[Asia/Jayapura]",
	"tz.jersey.europe": "[Europe/Jersey]",
	"tz.jerusalem.asia": "[Asia/Jerusalem]",
	"tz.johannesburg.africa": "[Africa/Johannesburg]",
	"tz.johnston.pacific": "[Pacific/Johnston]",
	"tz.juba.africa": "[Africa/Juba]",
	"tz.jujuy.america": "[America/Jujuy]",
	"tz.jujuy.argentina.america": "[America/Argentina/Jujuy]",
	"tz.juneau.america": "[America/Juneau]",
	"tz.kabul.asia": "[Asia/Kabul]",
	"tz.kaliningrad.europe": "[Europe/Kaliningrad]",
	"tz.kamchatka.asia": "[Asia/Kamchatka]",
	"tz.kampala.africa": "[Africa/Kampala]",
	"tz.kanton.pacific": "[Pacific/Kanton]",
	"tz.karachi.asia": "[Asia/Karachi]",
	"tz.kashgar.asia": "[Asia/Kashgar]",
	"tz.kathmandu.asia": "[Asia/Kathmandu]",
	"tz.katmandu.asia": "[Asia/Katmandu]",
	"tz.kerguelen.indian": "[Indian/Kerguelen]",
	"tz.khandyga.asia": "[Asia/Khandyga]",
	"tz.khartoum.africa": "[Africa/Khartoum]",
	"tz.kiev.europe": "[Europe/Kiev]",
	"tz.kigali.africa": "[Africa/Kigali]",
	"tz.kinshasa.africa": "[Africa/Kinshasa]",
	"tz.kiritimati.pacific": "[Pacific/Kiritimati]",
	"tz.kirov.europe": "[Europe/Kirov]",
	"tz.knox-in.america": "[America/Knox_IN]",
	"tz.knox.indiana.america": "[America/Indiana/Knox]",
	"tz.kolkata.asia": "[Asia/Kolkata]",
	"tz.kosrae.pacific": "[Pacific/Kosrae]",
	"tz.kralendijk.america": "[America/Kralendijk]",
	"tz.krasnoyarsk.asia": "[Asia/Krasnoyarsk]",
	"tz.kuala-lumpur.asia": "[Asia/Kuala_Lumpur]",
	"tz.kuching.asia": "[Asia/Kuching]",
	"tz.kuwait.asia": "[Asia/Kuwait]",
	"tz.kwajalein": "[Kwajalein]",
	"tz.kwajalein.pacific": "[Pacific/Kwajalein]",
	"tz.kyiv.europe": "[Europe/Kyiv]",
	"tz.la-paz.america": "[America/La_Paz]",
	"tz.la-rioja.argentina.america": "[America/Argentina/La_Rioja]",
	"tz.lagos.africa": "[Africa/Lagos]",
	"tz.lhi.australia": "[Australia/LHI]",
	"tz.libreville.africa": "[Africa/Libreville]",
	"tz.libya": "[Libya]",
	"tz.lima.america": "[America/Lima]",
	"tz.lindeman.australia": "[Australia/Lindeman]",
	"tz.lisbon.europe": "[Europe/Lisbon]",
	"tz.ljubljana.europe": "[Europe/Ljubljana]",
	"tz.lome.africa": "[Africa/Lome]",
	"tz.london.europe": "[Europe/London]",
	"tz.longyearbyen.arctic": "[Arctic/Longyearbyen]",
	"tz.lord-howe.australia": "[Australia/Lord_Howe]",
	"tz.los-angeles.america": "[America/Los_Angeles]",
	"tz.louisville.america": "[America/Louisville]",
	"tz.louisville.kentucky.america": "[America/Kentucky/Louisville]",
	"tz.lower-princes.america": "[America/Lower_Princes]",
	"tz.luanda.africa": "[Africa/Luanda]",
	"tz.lubumbashi.africa": "[Africa/Lubumbashi]",
	"tz.lusaka.africa": "[Africa/Lusaka]",
	"tz.luxembourg.europe": "[Europe/Luxembourg]",
	"tz.macao.asia": "[Asia/Macao]",
	"tz.macau.asia": "[Asia/Macau]",
	"tz.maceio.america": "[America/Maceio]",
	"tz.macquarie.antarctica": "[Antarctica/Macquarie]",
	"tz.madeira.atlantic": "[Atlantic/Madeira]",
	"tz.madrid.europe": "[Europe/Madrid]",
	"tz.magadan.asia": "[Asia/Magadan]",
	"tz.mahe.indian": "[Indian/Mahe]",
	"tz.majuro.pacific": "[Pacific/Majuro]",
	"tz.makassar.asia": "[Asia/Makassar]",
	"tz.malabo.africa": "[Africa/Malabo]",
	"tz.maldives.indian": "[Indian/Maldives]",
	"tz.malta.europe": "[Europe/Malta]",
	"tz.managua.america": "[America/Managua]",
	"tz.manaus.america": "[America/Manaus]",
	"tz.manila.asia": "[Asia/Manila]",
	"tz.maputo.africa": "[Africa/Maputo]",
	"tz.marengo.indiana.america": "[America/Indiana/Marengo]",
	"tz.mariehamn.europe": "[Europe/Mariehamn]",
	"tz.marigot.america": "[America/Marigot]",
	"tz.marquesas.pacific": "[Pacific/Marquesas]",
	"tz.martinique.america": "[America/Martinique]",
	"tz.maseru.africa": "[Africa/Maseru]",
	"tz.matamoros.america": "[America/Matamoros]",
	"tz.mauritius.indian": "[Indian/Mauritius]",
	"tz.mawson.antarctica": "[Antarctica/Mawson]",
	"tz.mayotte.indian": "[Indian/Mayotte]",
	"tz.mazatlan.america": "[America/Mazatlan]",
	"tz.mbabane.africa": "[Africa/Mbabane]",
	"tz.mcmurdo.antarctica": "[Antarctica/McMurdo]",
	"tz.melbourne.australia": "[Australia/Melbourne]",
	"tz.mendoza.america": "[America/Mendoza]",
	"tz.mendoza.argentina.america": "[America/Argentina/Mendoza]",
	"tz.menominee.america": "[America/Menominee]",
	"tz.merida.america": "[America/Merida]",
	"tz.met": "[MET]",
	"tz.metlakatla.america": "[America/Metlakatla]",
	"tz.mexico-city.america": "[America/Mexico_City]",
	"tz.mexico.general": "[Mexico/General]",
	"tz.michigan.us": "[US/Michigan]",
	"tz.midway.pacific": "[Pacific/Midway]",
	"tz.minsk.europe": "[Europe/Minsk]",
	"tz.miquelon.america": "[America/Miquelon]",
	"tz.mogadishu.africa": "[Africa/Mogadishu]",
	"tz.monaco.europe": "[Europe/Monaco]",
	"tz.moncton.america": "[America/Moncton]",
	"tz.monrovia.africa": "[Africa/Monrovia]",
	"tz.monterrey.america": "[America/Monterrey]",
	"tz.montevideo.america": "[America/Montevideo]",
	"tz.monticello.kentucky.america": "[America/Kentucky/Monticello]",
	"tz.montreal.america": "[America/Montreal]",
	"tz.montserrat.america": "[America/Montserrat]",
	"tz.moscow.europe": "[Europe/Moscow]",
	"tz.mountain.canada": "[Canada/Mountain]",
	"tz.mountain.us": "[US/Mountain]",
	"tz.mst": "[MST]",
	"tz.mst7mdt": "[MST7MDT]",
	"tz.muscat.asia": "[Asia/Muscat]",
	"tz.nairobi.africa": "[Africa/Nairobi]",
	"tz.nassau.america": "[America/Nassau]",
	"tz.nauru.pacific": "[Pacific/Nauru]",
	"tz.navajo": "[Navajo]",
	"tz.ndjamena.africa": "[Africa/Ndjamena]",
	"tz.new-salem.north-dakota.america": "[America/North_Dakota/New_Salem]",
	"tz.new-york.america": "[America/New_York]",
	"tz.newfoundland.canada": "[Canada/Newfoundland]",
	"tz.niamey.africa": "[Africa/Niamey]",
	"tz.nicosia.asia": "[Asia/Nicosia]",
	"tz.nicosia.europe": "[Europe/Nicosia]",
	"tz.nipigon.america": "[America/Nipigon]",
	"tz.niue.pacific": "[Pacific/Niue]",
	"tz.nome.america": "[America/Nome]",
	"tz.norfolk.pacific": "[Pacific/Norfolk]",
	"tz.noronha.america": "[America/Noronha]",
	"tz.north-dakota.center.america": "[America/North_Dakota/Center]",
	"tz.north.australia": "[Australia/North]",
	"tz.nouakchott.africa": "[Africa/Nouakchott]",
	"tz.noumea.pacific": "[Pacific/Noumea]",
	"tz.novokuznetsk.asia": "[Asia/Novokuznetsk]",
	"tz.novosibirsk.asia": "[Asia/Novosibirsk]",
	"tz.nsw.australia": "[Australia/NSW]",
	"tz.nuuk.america": "[America/Nuuk]",
	"tz.nz": "[NZ]",
	"tz.nz-chat": "[NZ-CHAT]",
	"tz.ojinaga.america": "[America/Ojinaga]",
	"tz.omsk.asia": "[Asia/Omsk]",
	"tz.oral.asia": "[Asia/Oral]",
	"tz.oslo.europe": "[Europe/Oslo]",
	"tz.ouagadougou.africa": "[Africa/Ouagadougou]",
	"tz.pacific.canada": "[Canada/Pacific]",
	"tz.pacific.us": "[US/Pacific]",
	"tz.pago-pago.pacific": "[Pacific/Pago_Pago]",
	"tz.palau.pacific": "[Pacific/Palau]",
	"tz.palmer.antarctica": "[Antarctica/Palmer]",
	"tz.panama.america": "[America/Panama]",
	"tz.pangnirtung.america": "[America/Pangnirtung]",
	"tz.paramaribo.america": "[America/Paramaribo]",
	"tz.paris.europe": "[Europe/Paris]",
	"tz.perth.australia": "[Australia/Perth]",
	"tz.petersburg.indiana.america": "[America/Indiana/Petersburg]",
	"tz.phnom-penh.asia": "[Asia/Phnom_Penh]",
	"tz.phoenix.america": "[America/Phoenix]",
	"tz.pitcairn.pacific": "[Pacific/Pitcairn]",
	"tz.podgorica.europe": "[Europe/Podgorica]",
	"tz.pohnpei.pacific": "[Pacific/Pohnpei]",
	"tz.poland": "[Poland]",
	"tz.ponape.pacific": "[Pacific/Ponape]",
	"tz.pontianak.asia": "[Asia/Pontianak]",
	"tz.port-au-prince.america": "[America/Port-au-Prince]",
	"tz.port-moresby.pacific": "[Pacific/Port_Moresby]",
	"tz.port-of-spain.america": "[America/Port_of_Spain]",
	"tz.porto-acre.america": "[America/Porto_Acre]",
	"tz.porto-novo.africa": "[Africa/Porto-Novo]",
	"tz.porto-velho.america": "[America/Porto_Velho]",
	"tz.portugal": "[Portugal]",
	"tz.prague.europe": "[Europe/Prague]",
	"tz.prc": "[PRC]",
	"tz.pst8pdt": "[PST8PDT]",
	"tz.puerto-rico.america": "[America/Puerto_Rico]",
	"tz.punta-arenas.america": "[America/Punta_Arenas]",
	"tz.pyongyang.asia": "[Asia/Pyongyang]",
	"tz.qatar.asia": "[Asia/Qatar]",
	"tz.qostanay.asia": "[Asia/Qostanay]",
	"tz.queensland.australia": "[Australia/Queensland]",
	"tz.qyzylorda.asia": "[Asia/Qyzylorda]",
	"tz.rainy-river.america": "[America/Rainy_River]",
	"tz.rangoon.asia": "[Asia/Rangoon]",
	"tz.rankin-inlet.america": "[America/Rankin_Inlet]",
	"tz.rarotonga.pacific": "[Pacific/Rarotonga]",
	"tz.recife.america": "[America/Recife]",
	"tz.regina.america": "[America/Regina]",
	"tz.resolute.america": "[America/Resolute]",
	"tz.reunion.indian": "[Indian/Reunion]",
	"tz.reykjavik.atlantic": "[Atlantic/Reykjavik]",
	"tz.riga.europe": "[Europe/Riga]",
	"tz.rio-branco.america": "[America/Rio_Branco]",
	"tz.rio-gallegos.argentina.america": "[America/Argentina/Rio_Gallegos]",
	"tz.riyadh.asia": "[Asia/Riyadh]",
	"tz.roc": "[ROC]",
	"tz.rok": "[ROK]",
	"tz.rome.europe": "[Europe/Rome]",
	"tz.rosario.america": "[America/Rosario]",
	"tz.rothera.antarctica": "[Antarctica/Rothera]",
	"tz.saigon.asia": "[Asia/Saigon]",
	"tz.saipan.pacific": "[Pacific/Saipan]",
	"tz.sakhalin.asia": "[Asia/Sakhalin]",
	"tz.salta.argentina.america": "[America/Argentina/Salta]",
	"tz.samara.europe": "[Europe/Samara]",
	"tz.samarkand.asia": "[Asia/Samarkand]",
	"tz.samoa.pacific": "[Pacific/Samoa]",
	"tz.samoa.us": "[US/Samoa]",
	"tz.san-juan.argentina.america": "[America/Argentina/San_Juan]",
	"tz.san-luis.argentina.america": "[America/Argentina/San_Luis]",
	"tz.san-marino.europe": "[Europe/San_Marino]",
	"tz.santa-isabel.america": "[America/Santa_Isabel]",
	"tz.santarem.america": "[America/Santarem]",
	"tz.santiago.america": "[America/Santiago]",
	"tz.santo-domingo.america": "[America/Santo_Domingo]",
	"tz.sao-paulo.america": "[America/Sao_Paulo]",
	"tz.sao-tome.africa": "[Africa/Sao_Tome]",
	"tz.sarajevo.europe": "[Europe/Sarajevo]",
	"tz.saratov.europe": "[Europe/Saratov]",
	"tz.saskatchewan.canada": "[Canada/Saskatchewan]",
	"tz.scoresbysund.america": "[America/Scoresbysund]",
	"tz.seoul.asia": "[Asia/Seoul]",
	"tz.shanghai.asia": "[Asia/Shanghai]",
	"tz.shiprock.america": "[America/Shiprock]",
	"tz.simferopol.europe": "[Europe/Simferopol]",
	"tz.singapore": "[Singapore]",
	"tz.singapore.asia": "[Asia/Singapore]",
	"tz.sitka.america": "[America/Sitka]",
	"tz.skopje.europe": "[Europe/Skopje]",
	"tz.sofia.europe": "[Europe/Sofia]",
	"tz.south-georgia.atlantic": "[Atlantic/South_Georgia]",
	"tz.south-pole.antarctica": "[Antarctica/South_Pole]",
	"tz.south.australia": "[Australia/South]",
	"tz.srednekolymsk.asia": "[Asia/Srednekolymsk]",
	"tz.st-barthelemy.america": "[America/St_Barthelemy]",
	"tz.st-helena.atlantic": "[Atlantic/St_Helena]",
	"tz.st-johns.america": "[America/St_Johns]",
	"tz.st-kitts.america": "[America/St_Kitts]",
	"tz.st-lucia.america": "[America/St_Lucia]",
	"tz.st-thomas.america": "[America/St_Thomas]",
	"tz.st-vincent.america": "[America/St_Vincent]",
	"tz.stanley.atlantic": "[Atlantic/Stanley]",
	"tz.stockholm.europe": "[Europe/Stockholm]",
	"tz.swift-current.america": "[America/Swift_Current]",
	"tz.sydney.australia": "[Australia/Sydney]",
	"tz.syowa.antarctica": "[Antarctica/Syowa]",
	"tz.tahiti.pacific": "[Pacific/Tahiti]",
	"tz.taipei.asia": "[Asia/Taipei]",
	"tz.tallinn.europe": "[Europe/Tallinn]",
	"tz.tarawa.pacific": "[Pacific/Tarawa]",
	"tz.tashkent.asia": "[Asia/Tashkent]",
	"tz.tasmania.australia": "[Australia/Tasmania]",
	"tz.tbilisi.asia": "[Asia/Tbilisi]",
	"tz.tegucigalpa.america": "[America/Tegucigalpa]",
	"tz.tehran.asia": "[Asia/Tehran]",
	"tz.tel-aviv.asia": "[Asia/Tel_Aviv]",
	"tz.tell-city.indiana.america": "[America/Indiana/Tell_City]",
	"tz.thimbu.asia": "[Asia/Thimbu]",
	"tz.thimphu.asia": "[Asia/Thimphu]",
	"tz.thule.america": "[America/Thule]",
	"tz.thunder-bay.america": "[America/Thunder_Bay]",
	"tz.tijuana.america": "[America/Tijuana]",
	"tz.timbuktu.africa": "[Africa/Timbuktu]",
	"tz.tirane.europe": "[Europe/Tirane]",
	"tz.tiraspol.europe": "[Europe/Tiraspol]",
	"tz.tokyo.asia": "[Asia/Tokyo]",
	"tz.tomsk.asia": "[Asia/Tomsk]",
	"tz.tongatapu.pacific": "[Pacific/Tongatapu]",
	"tz.toronto.america": "[America/Toronto]",
	"tz.tortola.america": "[America/Tortola]",
	"tz.tripoli.africa": "[Africa/Tripoli]",
	"tz.troll.antarctica": "[Antarctica/Troll]",
	"tz.truk.pacific": "[Pacific/Truk]",
	"tz.tucuman.argentina.america": "[America/Argentina/Tucuman]",
	"tz.tunis.africa": "[Africa/Tunis]",
	"tz.turkey": "[Turkey]",
	"tz.uct": "[UCT]",
	"tz.ujung-pandang.asia": "[Asia/Ujung_Pandang]",
	"tz.ulaanbaatar.asia": "[Asia/Ulaanbaatar]",
	"tz.ulan-bator.asia": "[Asia/Ulan_Bator]",
	"tz.ulyanovsk.europe": "[Europe/Ulyanovsk]",
	"tz.universal": "[Universal]",
	"tz.urumqi.asia": "[Asia/Urumqi]",
	"tz.us.eastern": "[US/Eastern]",
	"tz.us.pacific": "[US/Pacific]",
	"tz.ushuaia.argentina.america": "[America/Argentina/Ushuaia]",
	"tz.ust-nera.asia": "[Asia/Ust-Nera]",
	"tz.utc": "[UTC]",
	"tz.uzhgorod.europe": "[Europe/Uzhgorod]",
	"tz.vaduz.europe": "[Europe/Vaduz]",
	"tz.vancouver.america": "[America/Vancouver]",
	"tz.vatican.europe": "[Europe/Vatican]",
	"tz.vevay.indiana.america": "[America/Indiana/Vevay]",
	"tz.victoria.australia": "[Australia/Victoria]",
	"tz.vienna.europe": "[Europe/Vienna]",
	"tz.vientiane.asia": "[Asia/Vientiane]",
	"tz.vilnius.europe": "[Europe/Vilnius]",
	"tz.vincennes.indiana.america": "[America/Indiana/Vincennes]",
	"tz.virgin.america": "[America/Virgin]",
	"tz.vladivostok.asia": "[Asia/Vladivostok]",
	"tz.volgograd.europe": "[Europe/Volgograd]",
	"tz.vostok.antarctica": "[Antarctica/Vostok]",
	"tz.w-su": "[W-SU]",
	"tz.wake.pacific": "[Pacific/Wake]",
	"tz.wallis.pacific": "[Pacific/Wallis]",
	"tz.warsaw.europe": "[Europe/Warsaw]",
	"tz.west.australia": "[Australia/West]",
	"tz.west.brazil": "[Brazil/West]",
	"tz.wet": "[WET]",
	"tz.whitehorse.america": "[America/Whitehorse]",
	"tz.winamac.indiana.america": "[America/Indiana/Winamac]",
	"tz.windhoek.africa": "[Africa/Windhoek]",
	"tz.winnipeg.america": "[America/Winnipeg]",
	"tz.yakutat.america": "[America/Yakutat]",
	"tz.yakutsk.asia": "[Asia/Yakutsk]",
	"tz.yancowinna.australia": "[Australia/Yancowinna]",
	"tz.yangon.asia": "[Asia/Yangon]",
	"tz.yap.pacific": "[Pacific/Yap]",
	"tz.yekaterinburg.asia": "[Asia/Yekaterinburg]",
	"tz.yellowknife.america": "[America/Yellowknife]",
	"tz.yerevan.asia": "[Asia/Yerevan]",
	"tz.yukon.canada": "[Canada/Yukon]",
	"tz.zagreb.europe": "[Europe/Zagreb]",
	"tz.zaporozhye.europe": "[Europe/Zaporozhye]",
	"tz.zulu": "[Zulu]",
	"tz.zurich.europe": "[Europe/Zurich]",
}