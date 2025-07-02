import {UserModel} from "$lib/models/classes/user.model";

export interface ISession {
	token:  string;
	user: 	UserModel;
	botId: 	string;
}