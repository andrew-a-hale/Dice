package dice.java;

import java.util.ArrayList;

public class NumericDiceSet implements DiceSet<NumericDie, Integer> {
    ArrayList<NumericDie> diceSet = new ArrayList<>();

    @Override
    public void add(NumericDie numericDie) {
        diceSet.add(numericDie);
    }

    public Integer throwDice() {
        return this.diceSet.stream()
                .map(NumericDie::roll)
                .peek(System.out::println)
                .reduce(0, Integer::sum);
    }
}
