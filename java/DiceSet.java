package dice.java;

public interface DiceSet<T, U> {
    void add(T t);

    U throwDice();
}

