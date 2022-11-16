import type { ICountry } from '../interfaces/icountry.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class CountryModel implements IDeserializable<ICountry>, ICountry {

    public id!:            string;
	public name!:          string;
    public slug!:          string;
	public bandName!:      string;
	public songName!:      string;
	public flag!:          string;
	public participating!: boolean;

	update() : CountryModel {
		let country: CountryModel = new CountryModel;

        fetch('http://localhost:8080/country',{
            method: "PUT",
            mode: 'cors',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(this)
        }).then(response => response.json())
        .then(data => {
			country = data
        })
        .catch((err) => {
           console.log(err)
        })

		return country
    };

	deserialize(input: ICountry): this {
		Object.assign(this, input);
		return this;
	}
}