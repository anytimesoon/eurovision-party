import type { ILogin } from '../interfaces/ilogin.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class LoginModel implements IDeserializable<ILogin>, ILogin {

    public token!:   string;
    public userId!: string;

	deserialize(input: ILogin): this {
		Object.assign(this, input);
		return this;
	}
}