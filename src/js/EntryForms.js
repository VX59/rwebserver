function LoadForms() {
    const EntryForm = document.createElement("form");
    EntryForm.name = "EntryForm";
    const PswdInput = document.createElement("input");
    PswdInput.type = "password";
    PswdInput.placeholder = "authentication key";
    const submit = document.createElement("submit");
    submit.name = "submit"

    body = document.body;

    EntryForm.appendChild(PswdInput);
    EntryForm.appendChild(submit);

    body.appendChild(EntryForm);
}

const root = 'http://localhost:8081';

export async function AudioPackageNavigator() {
    const pacakge_table = document.createElement("div");
    const file_table = document.createElement("div");
    file_table.id = "file-table";
    // get packages here
    const response = await fetch(root+'/packages')
    .then((response) => response.json())
    .then((data) => {
        for (const pkg_name of data.Packages) {
            const pkg = document.createElement("button");
            pkg.innerHTML = pkg_name;
            pacakge_table.appendChild(pkg);

            pkg.addEventListener("click", async function(){
                file_table.innerHTML = "";
                // get packages here
                const headers = new Headers({
                    'Content-Type': 'application/json',
                    'package-name': pkg_name
                })
                const options = {
                    method: 'GET',
                    headers: headers
                  };
                const response = await fetch(root+'/content/',options)
                .then((response) => response.json())
                .then((data) => {
                    for (const file_name of data.Packages) {
                        const file = document.createElement("p");
                        file.className = "audio-package-file";
                        file.innerHTML = file_name;
                        file_table.appendChild(file);
                    }
                });
                document.body.append(file_table);
            });

        }
    });
    document.body.append(pacakge_table);
}