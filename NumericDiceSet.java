package dice;

import java.util.ArrayList;

public class NumericDiceSet implements DiceSet<NumericDie, Integer> {
    ArrayList<NumericDie> dice = new ArrayList<>();
    @Override
    public void add(NumericDie numericDie) {
        dice.add(numericDie);
    }

    @Override
    public Integer throwDie() {
        return this.dice.stream()
            .map(NumericDie::roll)
            .peek(System.out::println)
            .reduce(0, Integer::sum);
    }
}
