package dice.java;

public enum FudgeFace {
    PLUS, MINUS, NOTHING;

    static int mapFudgeFaceToInt(FudgeFace face) {
        return switch (face) {
            case PLUS -> 1;
            case MINUS -> -1;
            case NOTHING -> 0;
        };
    }
}

