package dice;

public interface DiceSet<T, U> {
    void add(T t);
    U throwDie();
}

