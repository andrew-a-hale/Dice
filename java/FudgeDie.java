package dice.java;

import java.util.ArrayList;
import java.util.Random;

public class FudgeDie implements Die<FudgeFace> {
    int numberOfFaces = 6;
    ArrayList<FudgeFace> faces = new ArrayList<>();

    private FudgeDie() {
        this.faces.add(FudgeFace.PLUS);
        this.faces.add(FudgeFace.PLUS);
        this.faces.add(FudgeFace.MINUS);
        this.faces.add(FudgeFace.MINUS);
        this.faces.add(FudgeFace.NOTHING);
        this.faces.add(FudgeFace.NOTHING);
    }

    public static FudgeDie makeDie() {
        return new FudgeDie();
    }

    @Override
    public FudgeFace roll() {
        Random random = new Random();
        return this.faces.get(random.nextInt(numberOfFaces));
    }
}
