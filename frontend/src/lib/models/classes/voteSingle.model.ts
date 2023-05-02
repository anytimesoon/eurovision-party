export class VoteSingleModel {
    public userId!: string;
    public countrySlug!: string;
    public cat!: string;
    public score!: number;

    constructor(userId:string, countrySlug:string, cat:string, score:number) {
        this.userId = userId;
        this.countrySlug = countrySlug;
        this.cat = cat;
        this.score = score;
    }
}