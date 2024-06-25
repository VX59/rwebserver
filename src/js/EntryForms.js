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

export async function PageNavigator() {
    const table = document.createElement("ul");

    // get packages here
    const response = await fetch(root+'/packages')
    .then((response) => response.json())
    .then((data) => {
        for (const pkg_name of data.Packages) {
            const pkg = document.createElement("li");
            pkg.innerHTML = pkg_name;
            table.appendChild(pkg);
        }
    });
    document.body.append(table);
}