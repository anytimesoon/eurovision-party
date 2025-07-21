import type { ISession } from '../interfaces/isession.interface';
import type {UserModel} from "$lib/models/classes/user.model";

export class SessionModel implements ISession {
    public token!:  string;
    public user!:   UserModel;
    public botId!:  string;
}

