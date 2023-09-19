<script>
    import CreateModal from "../lib/components/CreateModal.svelte";

    let showCreateModal = false;
    let showCreatedModal = false;

    let link = "https://"

    let data = {}

    async function createMiniLink (targetUrl) {
		const res = await fetch('/a/mini/create', {
			method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
			body: JSON.stringify({
				url: targetUrl,
            })
		})
		
		const json = await res.json()
        data = json
        console.log(data)
        return data
	}

</script>

<section class="bg-gray-900 text-white">
  <p class="absolute bottom-10 left-1/2 transform -translate-x-1/2 text-center">
    This page is very buggy yes - I'm not a front-end developer.<br/>If you don't write a valid URL, nothing will happen.
  </p>
  <div
    class="mx-auto max-w-screen-xl px-4 py-32 lg:flex lg:h-screen lg:items-center"
  >
    <div class="mx-auto max-w-3xl text-center">
      <h1
        class="bg-gradient-to-r from-green-300 via-blue-500 to-purple-600 bg-clip-text text-3xl font-extrabold text-transparent sm:text-5xl"
      >
        miniLink - Simple link shortener.

        <span class="sm:block">Written with speed in mind.</span>
      </h1>

      <p class="mx-auto mt-4 max-w-xl sm:text-xl/relaxed">
        This project is just for show, and is not recommended for use. It might
        not always be up. Feel free to try it out though!
      </p>

      <div class="mt-8 flex flex-wrap justify-center gap-4">
        <button
          on:click={() => {showCreateModal = !showCreateModal}}
          class="block w-full rounded border border-blue-600 bg-blue-600 px-12 py-3 text-sm font-medium text-white hover:bg-transparent hover:text-white focus:outline-none focus:ring active:text-opacity-75 sm:w-auto"
          href="/get-started"
        >
          Make a link!
        </button>

        <a
          class="block w-full rounded border border-blue-600 px-12 py-3 text-sm font-medium text-white hover:bg-blue-600 focus:outline-none focus:ring active:bg-blue-500 sm:w-auto"
          href="https://github.com/barealek/miniLink"
          target="_blank"
        >
          Source Code
        </a>
      </div>
    </div>
  </div>
</section>


<CreateModal bind:showModal={showCreateModal}>

    
    <div class="relative inline-block px-4 pt-5 pb-4 overflow-hidden text-left align-bottom transition-all transform bg-white rounded-lg shadow-xl dark:bg-gray-900 sm:my-8 sm:w-full sm:max-w-sm sm:p-6 sm:align-middle">
        <h3 class="text-lg font-medium leading-6 text-gray-800 capitalize dark:text-white" id="modal-title">
            Create a new miniLink
        </h3>
        <p class="mt-2 text-sm text-gray-500 dark:text-gray-400">
            Fill out the field below to get a shortened link.
        </p>

        <form class="mt-4" action="#">
            <label for="target-url" class="text-sm text-gray-700 dark:text-gray-200">
                Target URL
            </label>

            <label class="block mt-3" for="target-url">
                <input type="target-url" name="target-url" id="target-url" placeholder="https://..." bind:value={link} class="block w-full px-4 py-3 text-sm text-gray-700 bg-white border border-gray-200 rounded-md focus:border-blue-400 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 dark:border-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:focus:border-blue-300" />
            </label>

            

            <div class="mt-4 sm:flex sm:items-center">
                <button on:click={() => {createMiniLink(link); showCreatedModal = true}} type="button" class="w-max px-4 py-2 mt-3 text-sm font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-md sm:mt-0 sm:w-1/2 sm:mx-2 hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40">
                    Create miniLink
                </button>
            </div>
        </form>
    </div>
</CreateModal>

<CreateModal bind:showModal={showCreatedModal}>
    <div class="h-80">
          <div x-cloak class="fixed inset-0 z-10 bg-gray-700/50"></div>
          <div x-cloak class="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-0">
            <div class="mx-auto overflow-hidden rounded-lg bg-gray-900 shadow-xl sm:w-full sm:max-w-xl">
              <div class="relative p-6">
                <div class="flex gap-4">
                  <div class="flex h-10 w-10 items-center justify-center rounded-full bg-green-900/50 text-green-500">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="h-6 w-6">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <div class="flex-1">
                    <h3 class="text-lg font-medium text-gray-100">miniLink created!</h3>
                    <div class="mt-2 text-sm text-gray-100">The miniLink was successfully created. Head over to {"https://minilink.imalek.me/a/" + data.minilink}</div>
                  </div>
                </div>
                <div class="mt-6 flex justify-end gap-3">
                  <button type="button" class="rounded-lg border border-primary-500 bg-primary-500 px-4 py-2 text-center text-sm font-medium text-white shadow-sm transition-all hover:border-primary-700 hover:bg-primary-700 focus:ring focus:ring-primary-200 disabled:cursor-not-allowed disabled:border-primary-300 disabled:bg-primary-300">Confirm</button>
                </div>
              </div>
            </div>
          </div>
      </div>
</CreateModal>