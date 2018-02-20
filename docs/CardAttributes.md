# Card Attributes
This is a live document containing various features and statistics we are planning to implement for our Cards. The ideas are subject to change, as our goals and thought-processes may evolve during the process.

Additionally, please refer to our [Battle Mechanics](../master/BattleMechanics.md) document to learn how these card attributes are used strategically in battle.


## Sprint 1 Brainstorming
### Dividing up Cards into Classes and Subclasses
1. **Physical**
	* Subclasses:
		* Tank
		* Weapon Master
		* Unarmed (hand-to-hand combat)
2. **Support**
	* Subclasses:
		* Healer
		* Buffer / Status Effect
3. **Uber**
	* Subclasses:
		* Mime (mimic attacks)
		* Evasive
		* Last Stand (hurts itself when attacking)
4. **Ranged**
	* Subclasses:
		* Caster
		* Archer
		* Firearms

---

### Aggro (aggression) Attribute

>NOTE: We are also considering a "Battle Formations" mechanic as an alternative to Aggro for adding complexity to how a user would assemble a party of cards to battle with. You can read about this feature in the [Battle Mechanics](../master/BattleMechanics.md) document.

* **Force enemy cards to target the card with the highest aggro stat instead of randomly choosing a card to attack each round.**
	* Useful for cases where your whole team is all DPS or healers and 1 tank with large HP but low/medium ATK.
	* Tank cards have a chance of spawning with a “Taunt” passive ability that increases Aggro of that card, resulting in it always being attacked.
	* Tank cards also have a ~50% chance of spawning with a Taunt move (if it doesn’t have the passive ability).
	* All non-Tank cards have the same default value for Aggro which is lower than Tanks' value.

---
### Card Stats Ideas
* Attack
* Health
* Speed
* Aggression (IF no Battle Formation mechanic)
* Randomly-generated weakness towards other classes/subclasses
* Randomly-generated strength against other classes/subclasses
* Passive Ability
* 2-4 randomly generated moves from a list of moves usable by that subclass.
