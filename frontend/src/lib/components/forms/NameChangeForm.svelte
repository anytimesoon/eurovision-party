<script lang="ts">
    import FormButton from "$lib/components/buttons/FormButton.svelte";
    import {formButtonState} from "$lib/models/enums/formButtonState.enum";
    import ContentSave from "svelte-material-icons/ContentSave.svelte";
    import {currentUser, userStore} from "$lib/stores/user.store";
    import {put} from "$lib/utils/genericFetch";
    import {UserModel} from "$lib/models/classes/user.model";
    import {userEP} from "$lib/models/enums/endpoints.enum";

    interface Props {
        formToggle: VoidFunction;
    }

    let { formToggle }: Props = $props();

    let formState = $state(formButtonState.ENABLED)
    let form: HTMLFormElement = $state()

    const submit = async (e: Event) => {
        e.preventDefault()
        formState = formButtonState.SENDING;

        const formData = new FormData(form);
        const newUser = new UserModel(
            $currentUser.id,
            formData.get("name") as string,
            $currentUser.slug,
            $currentUser.icon,
            $currentUser.authLvl
        )

        const updatedUser = await put<UserModel, UserModel>(userEP.UPDATE, newUser)
            .then(res => UserModel.deserialize(res))

        $currentUser = updatedUser
        $userStore.set(updatedUser.id, updatedUser)

        formToggle()
        formState = formButtonState.ENABLED
    }
</script>
<form onsubmit={e => submit(e)} bind:this={form}>
    <div class="w-fit mx-auto flex justify-center">
        <input class="mr-3" type="text" name="name" bind:value={$currentUser.name} placeholder="Change your name"/>
        <FormButton buttonState={formState}>
            <ContentSave size="1.4em" />
        </FormButton>
    </div>
</form>