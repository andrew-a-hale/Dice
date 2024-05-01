package dice.java;

public class Main {
    public static void main(String[] args) {
        NumericDiceSet d20s = new NumericDiceSet();
        d20s.add(NumericDie.makeDie(20));
        d20s.add(NumericDie.makeDie(20));
        d20s.add(NumericDie.makeDie(20));
        System.out.println(d20s.throwDice());

        FudgeDiceSet fudgeDiceSet = new FudgeDiceSet();
        fudgeDiceSet.add(FudgeDie.makeDie());
        fudgeDiceSet.add(FudgeDie.makeDie());
        fudgeDiceSet.add(FudgeDie.makeDie());
        System.out.println(fudgeDiceSet.throwDice());
    }
}

