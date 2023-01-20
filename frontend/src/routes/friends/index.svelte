<script lang="ts">
    import AllUsers from "$lib/components/AllUsers.svelte";
    import {sendCreateOrUpdate} from "../../lib/helpers/sender.helper";
    import {NewUserModel} from "../../lib/models/classes/user.model";
    import {auth} from "../../lib/models/enums/endpoints.enum";

    function formHandler(e:Event) {
        const formData = new FormData(e.target as HTMLFormElement)
        console.log(formData)
        let newUser: NewUserModel = new NewUserModel(<string>formData.get("name"), <string>formData.get("email"))

        // newUser.name = <string>formData.get("name")
        // newUser.email = <string>formData.get("email")
        let registeredUser:NewUserModel
        sendCreateOrUpdate<NewUserModel, NewUserModel>(auth.REGISTER, newUser, "POST").then( data => {
            registeredUser = data.body
            document.getElementById("registered").innerText = registeredUser.name
        })
    }
</script>

<form on:submit|preventDefault={formHandler}>
    name <input type="text" name="name" />
    email <input type="text" name="email" />
    <input type="submit">
</form>
<AllUsers />
<div id="registered"></div>