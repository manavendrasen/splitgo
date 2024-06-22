import API from "./api";

export const getCategories = {
  queryKey: ["categories"],
  queryFn: async () => {
    const res = await API.get("/category");
    return res.data;
  },
};
