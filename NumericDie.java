package dice;

import java.util.ArrayList;
import java.util.Random;

public class NumericDie implements Die<Integer> {
    Integer numberOfFaces;
    ArrayList<Integer> faces = new ArrayList<>();

    private NumericDie(Integer numberOfFaces) {
        this.numberOfFaces = numberOfFaces;
        for (int i = 0; i < numberOfFaces; i++) {
            this.faces.add(i + 1);
        }
    }

    public static NumericDie makeDie(Integer numberOfFaces) {
        return new NumericDie(numberOfFaces);
    }

    @Override
    public Integer roll() {
        Random random = new Random();
        return this.faces.get(random.nextInt(this.numberOfFaces));
    }
}
