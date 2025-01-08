import type { ICountry } from '../interfaces/icountry.interface';

export class CountryModel implements ICountry {

	public name!:          string;
    public slug!:          string;
	public bandName!:      string;
	public songName!:      string;
	public flag!:          string;
	public participating!: boolean;

	constructor(name:string,
				slug:string,
				bandName:string,
				songName:string,
				flag:string,
				participating:boolean) {
		this.slug = slug
		this.name = name
		this.bandName = bandName
		this.songName = songName
		this.flag = flag
		this.participating = participating
	}

	static deserialize(input: ICountry): ICountry {
		return new CountryModel(
			input.name,
			input.slug,
			input.bandName,
			input.songName,
			input.flag,
			input.participating
		);
	}
}