#!/usr/local/bin/python

import cassiopeia as cass
import os
import sys
import time

with open("riot_key.txt", 'r') as f:
    RIOT_API_KEY = f.read()
print(RIOT_API_KEY)
cass.set_riot_api_key(RIOT_API_KEY)

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

me = cass.Summoner(name="PistoisorBlades", region="NA")

matches = me.match_history

match = get_latest(matches, "Renekton")
red_champs = [p.champion.name for p in match.red_team.participants]
blue_champs = [p.champion.name for p in match.blue_team.participants]
win = ("Camille" in red_champs and match.red_team.win) or \
    ("Camille" in blue_champs and match.blue_team.win)
participant = [p for p in match.participants if p.champion.name == "Camille"][0]

print("KDA: {}".format(round(participant.stats.kda, 2)))
sys.stdout.flush
