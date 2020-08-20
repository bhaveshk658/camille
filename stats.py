#!/usr/local/bin/python

def get_latest_kda(matchup):
    import cassiopeia as cass
    import os
    import sys

    def check(match, matchup):
        red_champs = [p.champion.name for p in match.red_team.participants]
        blue_champs = [p.champion.name for p in match.blue_team.participants]

        return ("Camille" in blue_champs and matchup in red_champs) or  \
            ("Camille" in red_champs and matchup in blue_champs)

    def get_latest(matches, matchup):
        for match in matches:
            if check(match, matchup):
                os.dup2(save[0], 1)
                os.dup2(save[1], 2)
                # close the temporary fds
                os.close(null_fds[0])
                os.close(null_fds[1])
                print("Match found")
                return match

        print("No recent matches found")
        return None

    RIOT_API_KEY = "RGAPI-091c05d6-a248-47e0-b9fc-20e3913648b7"
    cass.set_riot_api_key(RIOT_API_KEY)

    # open 2 fds
    null_fds = [os.open(os.devnull, os.O_RDWR) for x in range(2)]
    # save the current file descriptors to a tuple
    save = os.dup(1), os.dup(2)
    # put /dev/null fds on 1 and 2
    os.dup2(null_fds[0], 1)
    os.dup2(null_fds[1], 2)

    me = cass.Summoner(name="PistoisorBlades", region="NA")

    matches = me.match_history
    
    match = get_latest(matches, "Garen")
    red_champs = [p.champion.name for p in match.red_team.participants]
    blue_champs = [p.champion.name for p in match.blue_team.participants]
    win = ("Camille" in red_champs and match.red_team.win) or \
        ("Camille" in blue_champs and match.blue_team.win)
    participant = [p for p in match.participants if p.champion.name == "Camille"][0]
    
    print("KDA: {}".format(participant.stats.kda))
    sys.stdout.flush()