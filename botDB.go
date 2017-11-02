// DO NOT EDIT THIS FILE. IT IS AUTO-GENERATED BY "cmd/gen-db.go". //
/*
 * bot: Bot list, (C) 2017 Zoe.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package surferua

func init() {
	botDBSize = 16
	botDB = []string {
		"Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)",
		"Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5 (like Gecko) (Exabot-Thumbnails)",
		"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
		"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
		"Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)",
		"Sogou Orion spider/3.0( http://www.sogou.com/docs/help/webmasters.htm#07)",
		"ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)",
		"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)",
		"facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		"Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
		"DuckDuckBot/1.1; (+http://duckduckgo.com/duckduckbot.html)",
		"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
		"msnbot/2.0b (+http://search.msn.com/msnbot.htm)",
		"Mozilla/5.0 (compatible; MJ12bot/v1.4.5; http://www.majestic12.co.uk/bot.php)",
		"SimplePie/1.3.1 (Feed Parser; http://simplepie.org; Allow like Gecko)",
	}
}


// Get full user-agent string of google.
func NewBotGoogle() string {
	return "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
}

// Get full user-agent string of blex.
func NewBotBlex() string {
	return "Mozilla/5.0 (compatible; BLEXBot/1.0; +http://webmeup-crawler.com/)"
}

// Get full user-agent string of scoutJet.
func NewBotScoutJet() string {
	return "Mozilla/5.0 (compatible; ScoutJet; +http://www.scoutjet.com/)"
}

// Get full user-agent string of sogouOrin.
func NewBotSogouOrin() string {
	return "Sogou Orion spider/3.0( http://www.sogou.com/docs/help/webmasters.htm#07)"
}

// Get full user-agent string of alexa.
func NewBotAlexa() string {
	return "ia_archiver (+http://www.alexa.com/site/help/webmasters; crawler@alexa.com)"
}

// Get full user-agent string of yahoo.
func NewBotYahoo() string {
	return "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)"
}

// Get full user-agent string of duck.
func NewBotDuck() string {
	return "DuckDuckBot/1.1; (+http://duckduckgo.com/duckduckbot.html)"
}

// Get full user-agent string of facebook.
func NewBotFacebook() string {
	return "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)"
}

// Get full user-agent string of mj12.
func NewBotMj12() string {
	return "Mozilla/5.0 (compatible; MJ12bot/v1.4.5; http://www.majestic12.co.uk/bot.php)"
}

// Get full user-agent string of simplePie.
func NewBotSimplePie() string {
	return "SimplePie/1.3.1 (Feed Parser; http://simplepie.org; Allow like Gecko)"
}

// Get full user-agent string of yandex.
func NewBotYandex() string {
	return "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)"
}

// Get full user-agent string of msn.
func NewBotMsn() string {
	return "msnbot/2.0b (+http://search.msn.com/msnbot.htm)"
}

// Get full user-agent string of bing.
func NewBotBing() string {
	return "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)"
}

// Get full user-agent string of baidu.
func NewBotBaidu() string {
	return "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
}

// Get full user-agent string of sogou.
func NewBotSogou() string {
	return "Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)"
}

// Get full user-agent string of exa.
func NewBotExa() string {
	return "Mozilla/5.0 (compatible; Konqueror/3.5; Linux) KHTML/3.5.5 (like Gecko) (Exabot-Thumbnails)"
}

