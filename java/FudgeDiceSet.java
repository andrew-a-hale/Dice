package dice.java;

import java.util.ArrayList;

public class FudgeDiceSet implements DiceSet<FudgeDie, Integer> {
    ArrayList<FudgeDie> dice = new ArrayList<>();

    @Override
    public void add(FudgeDie fudgeDie) {
        this.dice.add(fudgeDie);
    }

    @Override
    public Integer throwDice() {
        return this.dice.stream()
                .map(FudgeDie::roll)
                .peek(System.out::println)
                .map(FudgeFace::mapFudgeFaceToInt)
                .reduce(0, Integer::sum);
    }
}
