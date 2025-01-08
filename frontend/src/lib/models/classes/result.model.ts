import {voteCats} from "$lib/models/enums/categories.enum";
import type {IResultModel} from "$lib/models/interfaces/iresultmodel.interface";

export class ResultModel implements IResultModel {

    public countrySlug!:  string;
    public costume!:      number;
    public song!:         number;
    public performance!:  number;
    public props!:        number;
    public total!:        number;

    constructor(
        countrySlug: string,
        costume: number,
        song: number,
        performance: number,
        props: number,
        total: number
    ) {
        this.countrySlug = countrySlug
        this.costume = costume
        this.song = song
        this.performance = performance
        this.props = props
        this.total = total
    }

    getScore(cat: string): number {
        switch (cat) {
            case voteCats.COSTUME:
                return this.costume
            case voteCats.SONG:
                return this.song
            case voteCats.PERFORMANCE:
                return this.performance
            case voteCats.PROPS:
                return this.props
            default:
                return this.total
        }
    }

    static deserialize(input: IResultModel):IResultModel {
        return new ResultModel(
            input.countrySlug,
            input.costume,
            input.song,
            input.performance,
            input.props,
            input.total
        )
    }
}