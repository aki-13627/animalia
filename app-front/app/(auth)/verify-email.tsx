import React from "react";
import {
  View,
  Text,
  Button,
  StyleSheet,
  Keyboard,
  TouchableWithoutFeedback,
} from "react-native";
import { useForm, Controller } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useColorScheme } from "@/hooks/useColorScheme";
import { Colors } from "@/constants/Colors";
import { FormInput } from "@/components/FormInput";
import { useVerifyEmailScreen } from "@/features/auth/useVerifyEmailScreen";

const VerifyEmailInputSchema = z.object({
  email: z.string().email({ message: "無効なメールアドレス" }),
  code: z.string().min(1, { message: "確認コードを入力してください" }),
});

type VerifyEmailInput = z.infer<typeof VerifyEmailInputSchema>;

export default function VerifyEmailScreen() {
  const { email } = useLocalSearchParams<{ email?: string }>();
  const {
    control,
    handleSubmit,
    formState: { errors, isSubmitting },
  } = useForm<VerifyEmailInput>({
    resolver: zodResolver(VerifyEmailInputSchema),
    defaultValues: { email: email || "", code: "" },
  });
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { onSubmit } = useVerifyEmailScreen();

  return (
    <TouchableWithoutFeedback onPress={() => Keyboard.dismiss()}>
      <View style={styles.container}>
        <Text
          style={[styles.title, { color: Colors[colorScheme ?? "light"].text }]}
        >
          メール認証
        </Text>
        <Controller
          control={control}
          name="email"
          render={({ field: { onChange, value } }) => (
            <FormInput
              label="Email"
              value={value}
              onChangeText={onChange}
              theme={theme}
              secureTextEntry
              autoCapitalize="none"
              error={errors.email?.message}
            />
          )}
        />
        <Controller
          control={control}
          name="code"
          render={({ field: { onChange, value } }) => (
            <FormInput
              label="Code"
              value={value}
              onChangeText={onChange}
              theme={theme}
              autoCapitalize="none"
              error={errors.code?.message}
            />
          )}
        />
        <Button
          title={isSubmitting ? "処理中..." : "認証する"}
          onPress={handleSubmit(onSubmit)}
          disabled={isSubmitting}
          color={Colors[colorScheme ?? "light"].tint}
        />
      </View>
    </TouchableWithoutFeedback>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 16,
    justifyContent: "center",
  },
  title: {
    fontSize: 24,
    fontWeight: "bold",
    marginBottom: 16,
    textAlign: "center",
  },
  input: {
    height: 40,
    borderColor: "#ccc",
    borderRadius: 8,
    borderWidth: 1,
    marginBottom: 12,
    paddingHorizontal: 8,
  },
  error: {
    color: "red",
    marginBottom: 8,
  },
});
