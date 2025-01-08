import {voteCats} from "$lib/models/enums/categories.enum";

export class ResultPageStateModel {
    public userId!:     string;
    public category!:   string;

    constructor(
        userId: string = "",
        category: string = voteCats.TOTAL
    ) {
        this.userId = userId
        this.category = category
    }

    reset() {
        this.userId = ""
        this.category = voteCats.TOTAL
    }
}