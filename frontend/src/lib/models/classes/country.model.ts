import type { ICountry } from '../interfaces/icountry.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class CountryModel implements IDeserializable<ICountry>, ICountry {

	public name!:          string;
    public slug!:          string;
	public bandName!:      string;
	public songName!:      string;
	public flag!:          string;
	public participating!: boolean;

	deserialize(input: ICountry): this {
		Object.assign(this, input);
		return this;
	}
}