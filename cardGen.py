# This script is unused (for now) because I'm swapping to Solidity instead

# Constants used in algorithm
STAT_MIN = 1
STAT_MAX = 1000

class Card(object):
    classType = ""
    name = ""
    attack = 0
    health = 0
    speed = 0
    aggro = 0
    weakAgainst = ""
    strongAgainst = ""
    passive = ""
    moveOne = ""
    moveTwo = ""

'''
Classes:
    Physical {
        Tank
        Weapons Master
        Unarmed
    }

    Support {
        Healer
        Stat Buffer
    }

    Uber {
        Mime
        Evasive
        Last Stand
    }

    Ranged {
        Caster
        Archer
        Firearms
    }

Attributes {
    Attack
    Health
    Speed
    Aggression
    WeakAgainst
    StrongAgainst
    PassiveAbility
    Move1
    Move2
}


Base Attribues Based on Classes:
    Physical:
    - Aggression:

'''

def generateCard():
