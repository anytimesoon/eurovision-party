import {UserModel} from "$lib/models/classes/user.model";
import {CookieOpts} from "$lib/models/classes/session.model";

export interface ISession {
	token:  string;
	user: 	UserModel;
	botId: 	string;
}