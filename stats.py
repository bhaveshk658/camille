import cassiopeia as cass

RIOT_API_KEY = "RGAPI-091c05d6-a248-47e0-b9fc-20e3913648b7"
cass.set_riot_api_key(RIOT_API_KEY)

me = cass.Summoner(name="PistoisorBlades", region="NA")

all_champions = cass.Champions(region="NA")
camille = all_champions["Camille"]


matches = me.match_history

def check(match, matchup):
    red_champs = [p.champion.name for p in match.red_team.participants]
    blue_champs = [p.champion.name for p in match.blue_team.participants]

    return ("Camille" in blue_champs and matchup in red_champs) or  \
           ("Camille" in red_champs and matchup in blue_champs)

def get_latest(matches, matchup):
    for match in matches:
        if check(match, matchup):
            print("Match found")
            return match

    print("No recent matches found")
    return None

get_latest(matches, "Garen")