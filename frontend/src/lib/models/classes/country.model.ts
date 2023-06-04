import type { ICountry } from '../interfaces/icountry.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class CountryModel implements IDeserializable<ICountry>, ICountry {

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

	deserialize(input: ICountry): this {
		Object.assign(this, input);
		return this;
	}
}