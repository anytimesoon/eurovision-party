import {voteCats} from "$lib/models/enums/categories.enum";

export class ResultPageStateModel {
    public userId!:     string;
    public category!:   string;
    public sortByDescending!: boolean;

    constructor(
        userId: string = "",
        category: string = voteCats.TOTAL,
        sortByDescending: boolean = false
    ) {
        this.userId = userId
        this.category = category
        this.sortByDescending = sortByDescending
    }

    reset() {
        this.userId = ""
        this.category = voteCats.TOTAL
        this.sortByDescending = false
    }
}