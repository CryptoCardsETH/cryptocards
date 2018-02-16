# Battle Mechanics

This is a live document containing various components we are planning to implement for our Battle mechanics. The ideas are subject to change, as our goals and thought-processes may evolve during the process.


## Sprint 1 Brainstorming

User Story 14:
`As a developer, I would like to figure out the components and rules of the games`

### Features We Like
*  **Each card has its own unique passive ability to assist itself and/or the other cards in its party**
	* For example:
		* Increase attack [or] defense by 10% for all party members
		* Heal itself [and/or] its party members for a portion of their health after each round
		* Randomly picks one attack per round from its party and gives 1.5x dmg multiplier

* **NO Accuracy or Critical Hit core attributes.**
	* These components can only exist in battle as passive abilities.
	* Additionally, these Crit Chance / Evade capabilities can only be applied to that card and no one else in its party.

* **“Uber” cards that are difficult or even useless to use in a party but yet extremely effective when paired strategically with certain types of cards.**
	* For example:
		* A card that has very low HP, ATK, SPEED but has an extremely high evasiveness passive ability. Difficult and annoying to kill, but it will die very easily once an attack or two actually lands.
		* A card that can only mimic the moves of the other cards in its party [or] the enemy’s moves
		* A card with very high ATK and below-average HP that injures itself after every attack.

* **Different “Classes” and "Subclasses" of cards which each have advantages and drawbacks in battle.**
	* Current Ideas:
		1. **Physical**
			*  Subclasses:
				1. Tank
				2. Weapon Master
				3. Unarmed (hand-to-hand combat)
		2. **Support**
			* Subclasses:
			1. Healer
			2. Buffer / Status Effect
		3. **Uber**
			* Subclasses:
				1. Mime (mimic attacks)
				2. Evasive
				3. Last Stand (hurts itself when attacking)
		4. **Ranged**
			* Subclasses:
				1. Caster
				2. Archer
				3. Firearms
