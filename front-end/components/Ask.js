import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { getQuestion } from "../requests/requests";

const jsonReponse = [
  {
    show_id: 0,
    show_type: "",
    title: "",
    director: "",
    casting: "",
    country: "",
    date_added: "",
    release_year: 0,
    rating: "",
    duration: "",
    listed_in: "",
  },
  {
    show_id: 0,
    show_type: "",
    title: "",
    director: "",
    casting: "",
    country: "",
    date_added: "",
    release_year: 0,
    rating: "",
    duration: "",
    listed_in: "",
  },
];

export default function Ask() {
  const queryClient = useQueryClient();
  const { isLoading, isError, isSuccess, data, error, refetch } = useQuery({
    queryKey: ["question"],
    queryFn: () => getQuestion(jsonReponse),
    refetchOnWindowFocus: false,
  });

//   const updateMovieQueryMutation = useMutation(getQuestion, {
//     onSuccess: () => {
//       console.log("Invalidating...");
//       queryClient.invalidateQueries(["question", jsonReponse]);
//     },
//   });

  function handleYes(event) {
    event.preventDefault();
    if (data.raw_name === "release_year") {
        jsonReponse[0][data.raw_name] = parseInt(data.answer)
    }
    else {jsonReponse[0][data.raw_name] = data.answer}
    // updateMovieQueryMutation.mutate(jsonReponse);
    // queryClient.invalidateQueries({queryKey: ["question"]});
    refetch()
  }

  function handleNo(event) {
    event.preventDefault();
    if (data.raw_name === "release_year") {
        jsonReponse[1][data.raw_name] = parseInt(data.answer)
    }
    else {jsonReponse[1][data.raw_name] = data.answer}
    // updateMovieQueryMutation.mutate(jsonReponse);
    // queryClient.invalidateQueries({queryKey: ["question"]});
    refetch()
  }

  function handleDontKnow(event) {
    event.preventDefault();
    // updateMovieQueryMutation.mutate(jsonReponse);
    // queryClient.invalidateQueries({queryKey: ["question"]});
    refetch()
  }

  return (
    <div>
      {isSuccess && (
        <div className="flex flex-col gap-2">
          <p className="text-center">{data.question}</p>
          <button onClick={handleYes}>Yes</button>
          <button onClick={handleDontKnow}>Dont know</button>
          <button onClick={handleNo}>Now</button>
        </div>
      )}
    </div>
  );
}
