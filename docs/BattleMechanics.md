# Battle Mechanics

This is a live document containing various components we are planning to implement for our Battle mechanics. The ideas are subject to change, as our goals and thought-processes may evolve during the process.

Additionally, please refer to our [Card Attributes](../master/CardAttributes.md) document for information about our current ideas regarding Card stats, strengths/weaknesses, etc.

## Sprint 1 Brainstorming

User Story 14:
`As a developer, I would like to figure out the components and rules of the games`

### Card-Specific Battle Features We Like:
*  **Each card has its own unique passive ability to assist itself and/or the other cards in its party.**
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


----------


### Core Battle Ideas
* **Battle Formations**
	>NOTE: We are also considering an "Aggro" mechanic as an alternative to Battle Formations to add strategy to how a user would assemble a party of cards to battle with. You can read about this feature in the [Card Attributes](../master/CardAttributes.md) document.
	* Allow the users to choose between various party setups to arrange their cards in.
		* Using a 4-card battle as an example, players can choose between:
			* 2 on front-lines, 2 behind.
			* 1 on front-lines, 3 behind.
			* All 4 cards on the front-lines.
	* Cards will only be able to attack the enemy cards on the front-lines and must destroy the front cards before being able to target the back cards.
	* Battle Formation could allow more complexity by strategically setting up your cards in various ways to get an advantage.
		* For example:
			* Placing your high DPS with low HP cards behind a tank card so the enemy has to burn down the tank’s high HP before it can get to the cards doing all the damage.
			*  “Ranged” class of cards has lower-than-average HP with above-average ATK. They could get a slight damage increase if you place them closer to the enemies as opposed to protecting them with a tank or card with higher HP.
				* This adds a sacrificial aspect of strategy, since it’s safer to put them behind high HP cards so they don’t die, but in some circumstances it might be useful having that bonus damage even if they might die quicker.
			* Putting a “Healer” class card in the back so it will remain alive long enough to heal the party.
