import type { ILogin } from '../interfaces/ilogin.interface';
import type { IDeserializable } from '../interfaces/ideserializable.interface';

export class LoginModel implements IDeserializable<ILogin>, ILogin {

    public token!:   string;
    public userId!: string;

	constructor(token:string, userId:string) {
		this.token = token
		this.userId = userId
	}

	deserialize(input: ILogin): this {
		Object.assign(this, input);
		return this;
	}
}