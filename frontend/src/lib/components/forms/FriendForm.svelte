<script lang="ts">
    import {formButtonState} from "$lib/models/enums/formButtonState.enum.js";
    import FormButton from "$lib/components/buttons/FormButton.svelte";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";
    import {NewUserModel} from "$lib/models/classes/newUser.model";
    import {authEP} from "$lib/models/enums/endpoints.enum";
    import {post} from "$lib/utils/genericFetch";
    import {newUserStore} from "$lib/stores/newUser.store";
    import {currentUser} from "$lib/stores/user.store";

    let form: HTMLFormElement = $state()
    let formState = $state(formButtonState.ENABLED)

    const submit = async (e: Event) => {
        e.preventDefault()
        formState = formButtonState.SENDING
        const fd = new FormData(form)
        const newUser = new NewUserModel(fd.get("name") as string, $currentUser.id)

        const createdUser: NewUserModel = await post<NewUserModel, NewUserModel>(authEP.REGISTER, newUser)
            .then(res => NewUserModel.deserialize(res))
        $newUserStore = [...$newUserStore, createdUser]

        form.reset()

        formState = formButtonState.ENABLED
    }
</script>

<form bind:this={form} onsubmit={e => submit(e)}>
    <div class="w-fit mx-auto flex justify-center">
        <input class="mr-3" id="new-user-name" type="text" name="name" placeholder="Name"/>
        <FormButton buttonState={formState}>
            <ContentSave size="1.4em" /> Save
        </FormButton>
    </div>
</form>