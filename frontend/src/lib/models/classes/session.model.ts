import type { ISession } from '../interfaces/isession.interface';
import type {UserModel} from "$lib/models/classes/user.model";
import type {sameSite} from "$lib/models/enums/cookie.enum";
import type {IDeserializable} from "$lib/models/interfaces/ideserializable.interface";

export class SessionModel implements ISession {
    public token!:  string;
    public user!:   UserModel;
    public botId!:  string;
}
