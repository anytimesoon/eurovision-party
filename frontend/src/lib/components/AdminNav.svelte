<script lang="ts">
    export let page:string
</script>

<ul class="p-3 flex justify-center">
    <li class="inline py-2 px-3 {page === 'settings' ? 'border-b-2 border-secondary' : ''}"><a href="/settings">Personal</a></li>
    <li class="inline py-2 px-3 {page === 'countries' ? 'border-b-2 border-secondary' : ''}"><a href="/admin/countries">Countries</a></li>
    <li class="inline py-2 px-3 {page === 'acts' ? 'border-b-2 border-secondary' : ''}"><a href="/admin/acts">Acts</a></li>
    <li class="inline py-2 px-3 {page === 'friends' ? 'border-b-2 border-secondary' : ''}"><a href="/admin/friends">Friends</a></li>
</ul>