## Декоратор (Decorator)

Используется для расширения функциональности объектов путем динамического добавления объекту новых возможностей. Сущность работы декоратора заключается в обёртывании готового объекта новым функционалом, при этом весь оригинальный интерфейс объекта остается доступным, путем передачи декоратором всех запросов обернутому объекту.

При такой структуре нам не важно является ли компонент декоратором или конкретной реализацией, так как интерфейс у них совпадает, и мы можем делать цепочки декораторов. Тем самым динамически менять состояние и поведение объекта.